package local_job_processor

import (
	"log"
	"workflows/internal/job_queue"
	"workflows/internal/workflows"
	"workflows/shared/job_manager"
	"workflows/shared/job_worker"
	"workflows/shared/nodes"

	"github.com/reactivex/rxgo/v2"
)

type LocalJobProcessor struct {
	eventStream *workflows.EventStream
	jobManager  *job_manager.JobManager
	jobQueue    *job_queue.JobQueue
}

func NewLocalJobProcessor(jobQueue *job_queue.JobQueue) (*LocalJobProcessor, error) {

	jobManager := job_manager.NewJobManager()

	processor := &LocalJobProcessor{
		jobManager: &jobManager,
		jobQueue:   jobQueue,
	}

	// register your workers
	injectWorker := job_worker.NewNodeJobWorker("Inject", nodes.ProcessInject)
	debugWorker := job_worker.NewNodeJobWorker("Debug", nodes.ProcessDebug)
	httpReqWorker := job_worker.NewNodeJobWorker("HTTP", nodes.ProcessHttpRequest)
	ParserWorker := job_worker.NewNodeJobWorker("Parser", nodes.ProcessParser)
	IfWorker := job_worker.NewNodeJobWorker("If", nodes.ProcessIf)

	processor.jobManager.AddWorker(injectWorker)
	processor.jobManager.AddWorker(debugWorker)
	processor.jobManager.AddWorker(httpReqWorker)
	processor.jobManager.AddWorker(ParserWorker)
	processor.jobManager.AddWorker(IfWorker)

	return processor, nil
}

func (processor *LocalJobProcessor) Register(eventStream *workflows.EventStream) {
	processor.eventStream = eventStream
	go processor.registerCommandsHandler(eventStream.CommandsObservable)
	go processor.registerEventsHandler(eventStream.EventsObservable)
}

func (processor *LocalJobProcessor) registerCommandsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(
		func(v interface{}) {
			command := v.(workflows.WorkflowCommand)
			switch command.Type {
			}
		},
		func(err error) {
			log.Printf("command failed: %s", err.Error())
		},
		func() {},
	)
}

func (processor *LocalJobProcessor) registerEventsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(
		func(v interface{}) {
			event := v.(workflows.WorkflowEvent)
			switch event.Type {
			case workflows.JobCreated:
				processor.jobCreated(event)
			}
		},
		func(err error) {
			log.Printf("event failed: %s", err.Error())
		},
		func() {},
	)
}

func (processor *LocalJobProcessor) jobCreated(event workflows.WorkflowEvent) {

	// log.Println("JobCreated")

	body := event.Body.(workflows.JobCreatedEventBody)
	job := body.Job

	if !processor.jobManager.CanExecute(job.Type) {
		return
	}

	// The job can be executed by an handler.
	// Try to lock the job
	locked, err := processor.jobQueue.LockJob(job.ID)
	if err != nil {
		log.Printf("failed to lock job: %s", err)
		return
	}

	if !locked {
		// another processor was faster
		return
	}

	// log.Println("JobLocked")

	// this processor successfully locked the job,
	// so it is allowed to execute it
	result, err := processor.jobManager.Execute(job)
	if err != nil {
		log.Printf("failed to execute job: %s", err.Error())
		return
	}

	jobCompletedEvent := workflows.NewJobCompletedEvent(workflows.JobCompletedEventBody{
		WorkflowID: job.WorkflowID,
		NodeID:     job.NodeID,
		JobId:      job.ID,
		Result:     result,
	})
	processor.eventStream.AddEvent(jobCompletedEvent)
}
