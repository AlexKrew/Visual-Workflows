package workflows

import (
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

func newWorkflowContainer(eventStream *EventStream, workflow *Workflow) *WorkflowContainer {
	container := WorkflowContainer{
		EventStream: eventStream,
		State:       LoadingContainer,
		Workflow:    workflow,
	}

	return &container
}

func WorkflowContainerFromStorage(eventStream *EventStream, workflowId WorkflowID) (*WorkflowContainer, error) {
	workflow, err := WorkflowFromFilesystem(workflowId)
	if err != nil {
		return nil, err
	}

	container := newWorkflowContainer(eventStream, &workflow)
	err = container.Init()
	return container, err
}

func (container *WorkflowContainer) ID() WorkflowID {
	return container.Workflow.ID
}

func (container *WorkflowContainer) Init() error {

	container.EventStream.AddEvent(NewWorkflowInstanceCreatedEvent(WorkflowInstanceCreatedEventBody{
		Workflow: *container.Workflow,
	}))

	err := container.initialize()
	if err != nil {
		log.Panicf("initialization failed: %s", err.Error())
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

func (container *WorkflowContainer) Stop() {
	container.CronJobs.stopCronJobs()
}

func (container *WorkflowContainer) PublishOutput(nodeId NodeID, output shared_entities.JobPayload) {
	for _, item := range output {

		node, exists := container.Workflow.NodeByID(nodeId)
		if !exists {
			log.Panicln("node does not exist", nodeId)
			return
		}

		portId, exists := node.PortByIdentifier(item.PortIdentifier, item.GroupID)
		if !exists {
			log.Panicf("port by ident `%s` and groupId `%s` does not exist", item.PortIdentifier, item.GroupID)
			return
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
			container.MessageCache.SetMessage(connPort, item.Value)

			// if port belongs to ui node
			// an event is created
			connNode, _ := container.Workflow.NodeByID(connPort.NodeID)
			connPort, _ := connNode.PortByID(connPort.PortID)

			if connNode.IsUINode {
				valueChangedEvent := NewDashboardValueChangedEvent(DashboardValueChangedEventBody{
					WorkflowID: container.ID(),
					ElementID:  connNode.ID,
					Field:      connPort.Identifier,
					Value:      item.Value,
				})
				container.EventStream.AddEvent(valueChangedEvent)
			}
		}

	}
}

func (container *WorkflowContainer) TriggerConnectedNodes(nodeId NodeID) {

	node, exists := container.Workflow.NodeByID(nodeId)
	if !exists {
		log.Panicln("Node does not exist", nodeId)
		return
	}

	triggerPortID, err := node.TriggerOutputPortID()
	if err != nil {
		log.Panicln(err)
		return
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

func (container *WorkflowContainer) InputForNodeId(nodeId NodeID) (shared_entities.JobPayload, bool) {
	if _, exists := container.Workflow.NodeByID(nodeId); !exists {
		return nil, false
	}

	return container.MessageCache.JobPayloadForNodeId(nodeId)
}
