package remote_job_processor

import (
	"log"
	gatewayserver "workflows/internal/gateway_server"
	"workflows/internal/job_queue"
	"workflows/internal/workflows"

	"github.com/reactivex/rxgo/v2"
)

type RemoteJobProcessor struct {
	eventStream   *workflows.EventStream
	gatewayServer *gatewayserver.GatewayServer
	jobQueue      *job_queue.JobQueue
}

func NewRemoteJobProcessor(jobQueue *job_queue.JobQueue, gatewayServer *gatewayserver.GatewayServer) (*RemoteJobProcessor, error) {
	return &RemoteJobProcessor{
		jobQueue:      jobQueue,
		gatewayServer: gatewayServer,
	}, nil
}

func (processor *RemoteJobProcessor) Register(eventStream *workflows.EventStream) {
	processor.eventStream = eventStream
	go processor.registerCommandsHandler(eventStream.CommandsObservable)
	go processor.registerEventsHandler(eventStream.EventsObservable)
}

func (processor *RemoteJobProcessor) registerCommandsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(
		func(v interface{}) {
			command := v.(workflows.WorkflowCommand)
			switch command.Type {
			}
		},
		func(err error) {
			//
		},
		func() {
			//
		},
	)
}

func (processor *RemoteJobProcessor) registerEventsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(
		func(v interface{}) {
			event := v.(workflows.WorkflowEvent)
			switch event.Type {
			case workflows.JobCreated:
				processor.jobCreated(event)
			}
		},
		func(err error) {
			//
		},
		func() {
			//
		},
	)
}

func (processor *RemoteJobProcessor) jobCreated(event workflows.WorkflowEvent) {
	body := event.Body.(workflows.JobCreatedEventBody)
	job := body.Job

	if !processor.gatewayServer.CanExecute(job.Type) {
		return
	}

	locked, err := processor.jobQueue.LockJob(job.ID)
	if err != nil {
		log.Printf("failed to lock job: %s", err)
		return
	}

	if !locked {
		// another process was faster
		return
	}

	processor.gatewayServer.Execute(job)
}
