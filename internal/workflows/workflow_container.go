package workflows

import (
	"fmt"
	"workflows/internal/utils"
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
	InstanceID    WorkflowContainerID
	EventStream   *EventStream
	Workflow      Workflow
	State         ContainerState
	MessageCache  *MessageCache
	MessageRouter *MessageRouter
}

func ConstructWorkflowContainer(eventStream *EventStream) WorkflowContainer {
	container := WorkflowContainer{
		InstanceID:  utils.GetNewUUID(),
		EventStream: eventStream,
		State:       LoadingContainer,
	}

	return container
}

func (container *WorkflowContainer) Run(workflow Workflow) error {
	container.Workflow = workflow

	container.EventStream.AddEvent(NewWorkflowInstanceCreatedEvent(WorkflowInstanceCreatedEventBody{
		Workflow:   workflow,
		InstanceID: container.InstanceID,
	}))

	err := container.initialize()
	if err != nil {
		return err
	}

	container.State = WorkflowStopped
	container.EventStream.AddEvent(NewWorkflowReadyEvent(WorkflowReadyEventBody{InstanceID: container.InstanceID}))

	return nil
}

func (container *WorkflowContainer) Start() {
	for _, node := range container.Workflow.Nodes {
		if node.Type == "Inject" {
			container.EventStream.AddCommand(NewCreateJobCommand(CreateJobCommandBody{WorkflowInstanceID: container.InstanceID, NodeID: node.ID}))
		}
	}
}

func (container *WorkflowContainer) PublishOutput(nodeId NodeID, output map[string]Message) {
	for portId, message := range output {

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
		fmt.Println("Trigger ", node.NodeID)
		container.EventStream.AddCommand(NewCreateJobCommand(CreateJobCommandBody{WorkflowInstanceID: container.InstanceID, NodeID: node.NodeID}))
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

	messageCache, err := ConstructMessageCache(&container.Workflow)
	if err != nil {
		return err
	}

	messageRouter, err := ConstructMessageRouter(&container.Workflow)
	if err != nil {
		return err
	}

	container.MessageCache = &messageCache
	container.MessageRouter = &messageRouter

	return nil
}

func (container *WorkflowContainer) InputForNodeId(nodeId NodeID) (map[NodeID]Message, bool) {
	if _, exists := container.Workflow.NodeByID(nodeId); !exists {
		return nil, false
	}

	return container.MessageCache.MessagesForNodeId(nodeId)
}
