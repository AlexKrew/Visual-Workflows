package main

import (
	"sync"
	gatewayserver "workflows/internal/gateway_server"
	"workflows/internal/job_queue"
	"workflows/internal/processors/local_job_processor"
	"workflows/internal/processors/remote_job_processor"
	"workflows/internal/processors/sysout_exporter"
	"workflows/internal/processors/workflow_processor"
	"workflows/internal/workflows"
	"workflows/webserver"
)

var wg sync.WaitGroup

func main() {

	wg.Add(5)

	eventStream := workflows.ConstructEventStream()
	jobQueue, _ := job_queue.NewJobQueue()

	// Register Processors
	registerSysoutExporter(eventStream, "./logs/log.jsonl")

	registerLocalJobProcessor(eventStream, jobQueue)

	registerRemoteJobProcessor(eventStream, jobQueue)

	wfProcessor := registerWorkflowProcessor(eventStream, jobQueue)
	go webserver.StartServer(eventStream, wfProcessor)

	wg.Wait()
}

func registerSysoutExporter(eventStream *workflows.EventStream, logfile string) *sysout_exporter.SysoutExporter {
	sysoutExporter, err := sysout_exporter.ConstructSysoutExporter(logfile)
	if err != nil {
		panic(err)
	}
	sysoutExporter.Register(eventStream)
	return sysoutExporter
}

func registerWorkflowProcessor(eventStream *workflows.EventStream, jobQueue *job_queue.JobQueue) *workflow_processor.WorkflowProcessor {
	workflowProcessor, err := workflow_processor.NewWorkflowProcessor(jobQueue)
	if err != nil {
		panic(err)
	}

	workflowProcessor.Register(eventStream)
	return workflowProcessor
}

func registerLocalJobProcessor(eventStream *workflows.EventStream, jobQueue *job_queue.JobQueue) *local_job_processor.LocalJobProcessor {
	jobProcessor, err := local_job_processor.NewLocalJobProcessor(jobQueue)
	if err != nil {
		panic(err)
	}

	jobProcessor.Register(eventStream)
	return jobProcessor
}

func registerRemoteJobProcessor(eventStream *workflows.EventStream, jobQueue *job_queue.JobQueue) (*remote_job_processor.RemoteJobProcessor, error) {
	gwServer, err := gatewayserver.StartGatewayServer(50051, eventStream)
	if err != nil {
		panic(err)
	}

	jobProcessor, err := remote_job_processor.NewRemoteJobProcessor(jobQueue, gwServer)
	if err != nil {
		panic(err)
	}

	jobProcessor.Register(eventStream)

	return jobProcessor, nil
}

func testCreateWorkflowInstance(eventStream *workflows.EventStream, id string) {
	eventStream.AddCommand(createTestCommand(id))
}

func createTestCommand(id string) workflows.WorkflowCommand {
	return workflows.NewCreateWorkflowInstanceCommand(workflows.CreateWorkflowInstanceCommandBody{
		WorkflowID: id,
	})
}
