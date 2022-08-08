package workflows

import "workflows/internal/utils"

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
