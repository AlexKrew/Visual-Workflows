package workflow_processor

import (
	"errors"
	"fmt"
	"workflows/internal/client"
	"workflows/internal/workflows"

	"github.com/reactivex/rxgo/v2"
)

type WorkflowProcessor struct {
	Containers map[workflows.WorkflowContainerID]workflows.WorkflowContainer

	EventStream *workflows.EventStream
}

func ConstructWorkflowProcessor() (*WorkflowProcessor, error) {

	containers := make(map[workflows.WorkflowContainerID]workflows.WorkflowContainer)

	return &WorkflowProcessor{
		Containers: containers,
	}, nil
}

func (processor *WorkflowProcessor) WorkflowByID(workflowId workflows.WorkflowID) (workflows.Workflow, bool) {
	for _, container := range processor.Containers {
		if container.Workflow.ID == workflowId {
			return container.Workflow, true
		}
	}

	return workflows.Workflow{}, false
}

func (processor *WorkflowProcessor) Register(eventStream *workflows.EventStream) {
	processor.EventStream = eventStream
	go processor.registerCommandsHandler(eventStream.CommandsObservable)
	go processor.registerEventsHandler(eventStream.EventsObservable)
}

func (processor *WorkflowProcessor) registerCommandsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		processor.handleCommand(i.(workflows.WorkflowCommand))

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}

func (processor *WorkflowProcessor) registerEventsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		processor.handleEvent(i.(workflows.WorkflowEvent))

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}

// ---- Commands ----

func (processor *WorkflowProcessor) handleCommand(command workflows.WorkflowCommand) error {
	switch command.Type {
	case workflows.CreateWorkflowInstance:
		return processor.createWorkflowInstance(command)
	case workflows.CreateJob:
		return processor.createJob(command)
	// case workflows.LockJob:
	// 	return processor.lockJob(command)
	case workflows.CompleteJob:
		return processor.completeJob(command)
	}

	return nil
}

func (processor *WorkflowProcessor) createWorkflowInstance(command workflows.WorkflowCommand) error {
	body := command.Body.(workflows.CreateWorkflowInstanceCommandBody)

	workflow, err := workflows.WorkflowFromFilesystem(body.WorkflowID)
	if err != nil {
		return err
	}

	instance := workflows.ConstructWorkflowContainer(processor.EventStream)
	err = instance.Run(workflow)
	if err != nil {
		return err
	}

	processor.Containers[instance.InstanceID] = instance

	return nil
}

func (processor *WorkflowProcessor) createJob(command workflows.WorkflowCommand) error {
	body := command.Body.(workflows.CreateJobCommandBody)

	container, ok := processor.Containers[body.WorkflowInstanceID]
	if !ok {
		return errors.New("no workflow with this instance id")
	}

	node, exists := container.Workflow.NodeByID(body.NodeID)
	if !exists {
		return errors.New("no node with this id")
	}

	input, exists := container.InputForNodeId(body.NodeID)
	if !exists {
		return errors.New("no input for node with this id")
	}

	job := workflows.NewJob(node.Type, input, body.NodeID)

	jobCreatedEvent := workflows.JobCreatedEventBody{
		WorkflowInstanceID: body.WorkflowInstanceID,
		Job:                job,
	}
	processor.EventStream.AddEvent(workflows.NewJobCreatedEvent(jobCreatedEvent))

	return nil
}

// func (processor *WorkflowProcessor) lockJob(command workflows.WorkflowCommand) error {
// 	return nil
// }

func (processor *WorkflowProcessor) completeJob(command workflows.WorkflowCommand) error {
	return nil
}

// ---- Events ----

func (processor *WorkflowProcessor) handleEvent(event workflows.WorkflowEvent) error {
	switch event.Type {
	case workflows.WorkflowReady:
		fmt.Println("Workflow is ready", event.Body)
		processor.workflowReady(event)
	case workflows.JobCompleted:
		fmt.Println("Job Completed in workflow")
		processor.jobCompleted(event)
	}

	return nil
}

func (processor *WorkflowProcessor) workflowReady(event workflows.WorkflowEvent) {
	body := event.Body.(workflows.WorkflowReadyEventBody)

	for _, container := range processor.Containers {
		if container.InstanceID == body.InstanceID {
			fmt.Println("Starting container", body.InstanceID)
			container.Start()
		}
	}
}

func (processor *WorkflowProcessor) jobCompleted(event workflows.WorkflowEvent) {
	body := event.Body.(workflows.JobCompletedEventBody)
	results := body.Result.(client.JobResults)

	fmt.Println("LOG EVENTS", results.Logs)
	for _, log := range results.Logs {
		processor.EventStream.AddEvent(workflows.NewDebugEvent(workflows.DebugEventBody{
			WorkflowInstanceID: body.WorkflowInstanceID,
			WorkflowID:         "LOGEVENT",
			Value:              log,
		}))
	}

	container := processor.Containers[body.WorkflowInstanceID]
	resultMessages := make(map[string]workflows.Message)
	for key, msg := range results.Output {
		resultMessages[key] = workflows.Message{
			DataType: workflows.MessageTypeFromString(msg.Datatype),
			Value:    msg.Value,
		}
	}

	container.PublishOutput(body.Job.NodeID, resultMessages)
	fmt.Println("Messages updated")
	container.TriggerConnectedNodes(body.Job.NodeID)
}
