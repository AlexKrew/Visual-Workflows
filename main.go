package main

import (
	"sync"
	"time"
	"workflows/internal/job_queue"
	"workflows/internal/processors/local_job_processor"
	"workflows/internal/processors/sysout_exporter"
	"workflows/internal/processors/workflow_processor"
	"workflows/internal/utils"
	"workflows/internal/workflows"
	"workflows/webserver"
)

var wg sync.WaitGroup

func main() {

	wg.Add(5)

	// workerClient, _ := client.NewClient()

	eventStream := workflows.ConstructEventStream()

	// // Register Processors
	registerSysoutExporter(eventStream, "./logs/log.jsonl")

	// // Mandatory: Workflow logic
	jobQueue, _ := job_queue.NewJobQueue()

	// go gatewayserver.StartGatewayServer(50051)

	registerLocalJobProcessor(eventStream, jobQueue)

	wfProcessor := registerWorkflowProcessor(eventStream, jobQueue)

	// // Test sysout-exporter
	// // go testSysoutExporter(eventStream)

	time.Sleep(2 * time.Second)
	go testCreateWorkflowInstance(eventStream, "3d48d394-08e4-4858-a936-4fc7201be0a2")

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

func testSysoutExporter(eventStream *workflows.EventStream) {
	eventStream.AddEvent(createTestEvent())
	time.Sleep(2 * time.Second)
	eventStream.AddEvent(createTestEvent())
	time.Sleep(2 * time.Second)
	eventStream.AddEvent(createTestEvent())
	time.Sleep(2 * time.Second)
	eventStream.AddEvent(createTestEvent())
}

func createTestEvent() workflows.WorkflowEvent {
	return workflows.NewWorkflowStartedEvent(workflows.WorkflowStartedEventBody{
		WorkflowID: utils.GetNewUUID(),
	})
}

func testCreateWorkflowInstance(eventStream *workflows.EventStream, id string) {
	eventStream.AddCommand(createTestCommand(id))
}

func createTestCommand(id string) workflows.WorkflowCommand {
	return workflows.NewCreateWorkflowInstanceCommand(workflows.CreateWorkflowInstanceCommandBody{
		WorkflowID: id,
	})
}
