package workflow_processor

import (
	"errors"
	"fmt"
	"log"
	"workflows/internal/job_queue"
	"workflows/internal/workflows"
	"workflows/shared/shared_entities"

	"github.com/reactivex/rxgo/v2"
)

type WorkflowProcessor struct {
	Containers map[workflows.WorkflowID]*workflows.WorkflowContainer

	EventStream *workflows.EventStream
	jobQueue    *job_queue.JobQueue
}

func NewWorkflowProcessor(jobQueue *job_queue.JobQueue) (*WorkflowProcessor, error) {

	containers := make(map[workflows.WorkflowContainerID]*workflows.WorkflowContainer)

	return &WorkflowProcessor{
		Containers: containers,
		jobQueue:   jobQueue,
	}, nil
}

func (processor *WorkflowProcessor) WorkflowByID(workflowId workflows.WorkflowID) (*workflows.Workflow, bool) {
	for id, container := range processor.Containers {
		if id == workflowId {
			return container.Workflow, true
		}
	}

	return nil, false
}

func (processor *WorkflowProcessor) CreateContainer(workflowId workflows.WorkflowID) error {
	container, err := workflows.WorkflowContainerFromStorage(processor.EventStream, workflowId)
	if err != nil {
		return err
	}

	processor.Containers[container.ID()] = container

	return nil
}

func (processor *WorkflowProcessor) StartWorkflow(workflowId workflows.WorkflowID) error {

	availableWorkflows := []string{}
	for workflowId, _ := range processor.Containers {
		availableWorkflows = append(availableWorkflows, workflowId)
	}

	container, exists := processor.Containers[workflowId]
	if !exists {
		return errors.New("failed to start workflow. workflow does not exist")
	}
	container.Start()

	return nil
}

func (processor *WorkflowProcessor) ReloadContainer(workflowId workflows.WorkflowID) error {
	container, exists := processor.Containers[workflowId]
	if exists {
		container.Stop()
	}
	processor.Containers[workflowId] = nil

	return processor.CreateContainer(workflowId)
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
	log.Printf("WFProcessor: Received command: %d", command.Type)

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

	instance, err := workflows.WorkflowContainerFromStorage(processor.EventStream, body.WorkflowID)
	if err != nil {
		log.Panicf("ERR: %s", err.Error())
		return err
	}

	processor.Containers[instance.ID()] = instance

	return nil
}

func (processor *WorkflowProcessor) createJob(command workflows.WorkflowCommand) error {
	body := command.Body.(workflows.CreateJobCommandBody)

	container, ok := processor.Containers[body.WorkflowID]
	if !ok {
		log.Panicln("ERR: no workflow with this instance id")
		return errors.New("no workflow with this instance id")
	}

	node, exists := container.Workflow.NodeByID(body.NodeID)
	if !exists {
		log.Panicln("ERR: no node with this id")
		return errors.New("no node with this id")
	}

	input, exists := container.InputForNodeId(body.NodeID)
	if !exists {
		log.Panicln("ERR: no input for node with this id")
		return errors.New("no input for node with this id")
	}

	job := shared_entities.NewJob(node.Type, input, body.NodeID, container.ID())
	added := processor.jobQueue.AddJob(job)

	if added {
		jobCreatedEvent := workflows.JobCreatedEventBody{
			WorkflowID: job.WorkflowID,
			Job:        job,
		}
		processor.EventStream.AddEvent(workflows.NewJobCreatedEvent(jobCreatedEvent))
	}

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
	log.Printf("WFProcessor: Received event: %d", event.Type)

	switch event.Type {
	case workflows.WorkflowReady:
		// processor.workflowReady(event)
	case workflows.JobCompleted:
		processor.jobCompleted(event)
	case workflows.CronTriggered:
		processor.cronTriggered(event)
	}

	return nil
}

// func (processor *WorkflowProcessor) workflowReady(event workflows.WorkflowEvent) {
// 	body := event.Body.(workflows.WorkflowReadyEventBody)

// 	for _, container := range processor.Containers {
// 		if container.ID() == body.WorkflowID {
// 			container.Start()
// 		}
// 	}
// }

func (processor *WorkflowProcessor) jobCompleted(event workflows.WorkflowEvent) {

	body := event.Body.(workflows.JobCompletedEventBody)
	results := body.Result

	for _, log := range results.Logs {
		debugEvent := workflows.NewDebugEvent(workflows.DebugEventBody{
			WorkflowID: body.WorkflowID,
			Value:      log,
		})
		processor.EventStream.AddEvent(debugEvent)
	}

	container, exists := processor.Containers[body.WorkflowID]
	if !exists {
		panic("workflow does not exist")
	}

	// resultMessages := make(map[string]shared_entities.WorkflowMessage)
	// for key, msg := range results.Output {
	// 	resultMessages[key] = shared_entities.WorkflowMessage{
	// 		DataType: msg.DataType,
	// 		Value:    msg.Value,
	// 	}
	// }
	output := results.Output

	container.PublishOutput(body.NodeID, output)
	if !results.DontTrigger {
		container.TriggerConnectedNodes(body.NodeID)
	}
}

func (processor *WorkflowProcessor) cronTriggered(event workflows.WorkflowEvent) {
	body := event.Body.(workflows.CronTriggeredBody)

	container, exists := processor.Containers[body.WorkflowID]
	if !exists {
		log.Fatalln("trigger for unknown workflow id:", body.WorkflowID)
		return
	}

	log.Println("Trigger node", body.NodeID)
	container.TriggerConnectedNodes(body.NodeID)
}
