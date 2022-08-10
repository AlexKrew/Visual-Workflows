package job_queue_processor

import (
	"fmt"
	"workflows/internal/client"
	"workflows/internal/workflows"

	"github.com/reactivex/rxgo/v2"
)

type JobQueueProcessor struct {
	jobQueue    *JobQueue
	eventStream *workflows.EventStream
}

func ConstructJobQueueProcessor() (*JobQueueProcessor, error) {

	jobQueue, err := ConstructJobQueue()
	if err != nil {
		return nil, err
	}

	jobQueueProcessor := JobQueueProcessor{
		jobQueue: jobQueue,
	}

	return &jobQueueProcessor, nil
}

func (jobQueueProcessor *JobQueueProcessor) AddClient(client *client.Client) {
	jobQueueProcessor.jobQueue.client = client
}

func (processor *JobQueueProcessor) Register(eventStream *workflows.EventStream) {
	processor.eventStream = eventStream
	go processor.registerCommandsHandler(eventStream.CommandsObservable)
	go processor.registerEventsHandler(eventStream.EventsObservable)
}

func (processor *JobQueueProcessor) registerCommandsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		command := i.(workflows.WorkflowCommand)
		switch command.Type {
		case workflows.CompleteJob:
			processor.completeJob(command)
		}

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}

func (processor *JobQueueProcessor) completeJob(command workflows.WorkflowCommand) {
	body := command.Body.(workflows.CompleteJobCommandBody)

	completedJob, exists := processor.jobQueue.RemoveJob(body.JobID)
	if !exists {
		panic("could not remove job from queue: job does not exist")
	}

	event := workflows.JobCompletedEventBody{
		WorkflowInstanceID: body.WorflowInstanceID,
		Job:                completedJob,
		Result:             body.Output,
	}
	processor.eventStream.AddEvent(workflows.NewJobCompletedEvent(event))
}

// Events

func (processor *JobQueueProcessor) registerEventsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		event := i.(workflows.WorkflowEvent)

		switch event.Type {
		case workflows.JobCreated:
			processor.jobCreated(event)
		}

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}

func (processor *JobQueueProcessor) jobCreated(event workflows.WorkflowEvent) {
	body := event.Body.(workflows.JobCreatedEventBody)
	processor.jobQueue.AddJob(body.Job)

	results, err := processor.jobQueue.ExecuteJob(body.Job.ID)
	if err != nil {
		panic(err)
	}

	completedEvent := workflows.JobCompletedEventBody{WorkflowInstanceID: body.WorkflowInstanceID, Job: body.Job, Result: results}
	processor.eventStream.AddEvent(workflows.NewJobCompletedEvent(completedEvent))
}