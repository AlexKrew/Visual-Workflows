package workflows

import (
	"fmt"
	"log"
	"workflows/internal/utils"
	"workflows/shared/shared_entities"
)

type WorkflowContainerID = utils.UUID
type ContainerState = int

const (
	LoadingContainer = iota
	StartingContainer
	RunningContainer
	StoppingContainer
	StoppedContainer
)

type WorkflowContainer struct {
	EventStream   *EventStream
	Workflow      *Workflow
	State         ContainerState
	MessageCache  *MessageCache
	MessageRouter *MessageRouter
	CronJobs      *CronJobManager
}

func NewWorkflowContainer(eventStream *EventStream, workflow *Workflow) *WorkflowContainer {
	container := WorkflowContainer{
		EventStream: eventStream,
		State:       LoadingContainer,
		Workflow:    workflow,
	}

	return &container
}

func WorkflowFromStorage(eventStream *EventStream, workflowId WorkflowID) (*WorkflowContainer, error) {
	workflow, err := WorkflowFromFilesystem(workflowId)
	if err != nil {
		return nil, err
	}

	return NewWorkflowContainer(eventStream, &workflow), nil
}

func (container *WorkflowContainer) ID() WorkflowID {
	return container.Workflow.ID
}

func (container *WorkflowContainer) Run(workflow *Workflow) error {
	container.Workflow = workflow

	container.EventStream.AddEvent(NewWorkflowInstanceCreatedEvent(WorkflowInstanceCreatedEventBody{
		Workflow: *workflow,
	}))

	err := container.initialize()
	if err != nil {
		return err
	}

	container.State = WorkflowStopped
	container.EventStream.AddEvent(NewWorkflowReadyEvent(WorkflowReadyEventBody{WorkflowID: container.ID()}))

	return nil
}

func (container *WorkflowContainer) Start() {

	for _, node := range container.Workflow.Nodes {
		if node.Type == "Inject" {

			createJobCommand := NewCreateJobCommand(
				CreateJobCommandBody{
					WorkflowID: container.ID(),
					NodeID:     node.ID,
				},
			)
			container.EventStream.AddCommand(createJobCommand)
		}

		if node.Type == "CronJob" {
			cronJob, err := NewCronJob(node)
			if err != nil {
				log.Fatalf("Failed to create cronjob: %s", err.Error())
			}

			container.CronJobs.addCronJob(cronJob)
		}
	}

	container.CronJobs.startCronJobs()
}

func (container *WorkflowContainer) PublishOutput(nodeId NodeID, output map[string]shared_entities.WorkflowMessage) {
	for portIdentifier, message := range output {

		node, exists := container.Workflow.NodeByID(nodeId)
		if !exists {
			panic("node does not exist")
		}

		portId, exists := node.PortByIdentifier(portIdentifier)
		if !exists {
			panic("port by ident does not exist")
		}

		addr := PortAddress{
			NodeID: nodeId,
			PortID: portId,
		}

		uAddr := addr.UniquePortID()
		connPorts, exists := container.MessageRouter.connectedPorts[uAddr]
		if !exists {
			// no ports connected
			return
		}

		for _, connPort := range connPorts {
			container.MessageCache.SetMessage(connPort, message)

			// if port belongs to dashboard node
			// an event is created
			connNode, _ := container.Workflow.NodeByID(connPort.NodeID)
			connPort, _ := connNode.PortByID(connPort.PortID)

			if node.IsDashboardNode {
				valueChangedEvent := NewDashboardValueChangedEvent(DashboardValueChangedEventBody{
					WorkflowID: container.ID(),
					ElementID:  connNode.ID,
					Field:      connPort.Identifier,
					Value:      message.Value,
				})
				container.EventStream.AddEvent(valueChangedEvent)
			}
		}

	}
}

func (container *WorkflowContainer) TriggerConnectedNodes(nodeId NodeID) {

	node, exists := container.Workflow.NodeByID(nodeId)
	if !exists {
		fmt.Println("Node does not exists", node)
		panic("panic")
	}

	triggerPortID, err := node.TriggerOutputPortID()
	if err != nil {
		panic(err)
	}

	triggerAddr := PortAddress{
		NodeID: nodeId,
		PortID: triggerPortID,
	}

	triggerNodes := container.MessageRouter.connectedPorts[triggerAddr.UniquePortID()]
	for _, node := range triggerNodes {
		container.EventStream.AddCommand(NewCreateJobCommand(CreateJobCommandBody{NodeID: node.NodeID, WorkflowID: container.ID()}))
	}
}

// func ConstructWorkflowInstance(eventStream *EventStream) (WorkflowContainer, error) {

// 	container := WorkflowContainer{
// 		InstanceID:  utils.GetNewUUID(),
// 		EventStream: eventStream,
// 		Workflow:    workflow,
// 		State:       LoadingContainer,
// 	}

// 	eventStream.AddEvent(NewWorkflowInstanceCreatedEvent(WorkflowInstanceCreatedEventBody{
// 		Workflow:   workflow,
// 		InstanceID: container.InstanceID,
// 	}))

// 	// TODO: Run init logic of all services

// 	err := container.initialize()
// 	if err != nil {
// 		return WorkflowContainer{}, err
// 	}

// 	container.State = WorkflowStopped
// 	eventStream.AddEvent(NewWorkflowReadyEvent(WorkflowReadyEventBody{InstanceID: container.InstanceID}))

// 	return container, nil
// }

func (container *WorkflowContainer) initialize() error {

	messageCache, err := ConstructMessageCache(container.Workflow)
	if err != nil {
		return err
	}

	messageRouter, err := ConstructMessageRouter(container.Workflow)
	if err != nil {
		return err
	}

	cronJobs := NewCronJobManager(container.EventStream, container.Workflow)

	container.MessageCache = &messageCache
	container.MessageRouter = &messageRouter
	container.CronJobs = cronJobs

	return nil
}

func (container *WorkflowContainer) InputForNodeId(nodeId NodeID) (map[NodeID]shared_entities.WorkflowMessage, bool) {
	if _, exists := container.Workflow.NodeByID(nodeId); !exists {
		return nil, false
	}

	return container.MessageCache.MessagesForNodeId(nodeId)
}
