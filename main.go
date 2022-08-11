package main

import (
	"sync"
	"time"
	gatewayserver "workflows/internal/gateway_server"
	"workflows/internal/processors/job_queue_processor"
	"workflows/internal/processors/sysout_exporter"
	"workflows/internal/processors/workflow_processor"
	"workflows/internal/utils"
	"workflows/internal/workflows"
)

var wg sync.WaitGroup

func main() {

	wg.Add(5)

	go gatewayserver.StartGatewayServer(50051)

	// workerClient, _ := client.NewClient()

	// eventStream := workflows.ConstructEventStream()

	// // Register Processors
	// registerSysoutExporter(eventStream, "./logs/log.jsonl")

	// // Mandatory: Workflow logic
	// wfProcessor := registerWorkflowProcessor(eventStream)
	// jobQueueProcessor := registerJobQueueProcessor(eventStream)
	// jobQueueProcessor.AddClient(&workerClient)

	// time.Sleep(1 * time.Second)

	// // Test sysout-exporter
	// // go testSysoutExporter(eventStream)
	// go testCreateWorkflowInstance(eventStream, "3d48d394-08e4-4858-a936-4fc7201be0a2")

	// go server.StartServer(eventStream, wfProcessor)

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

func registerWorkflowProcessor(eventStream *workflows.EventStream) *workflow_processor.WorkflowProcessor {
	workflowProcessor, err := workflow_processor.ConstructWorkflowProcessor()
	if err != nil {
		panic(err)
	}

	workflowProcessor.Register(eventStream)
	return workflowProcessor
}

func registerJobQueueProcessor(eventStream *workflows.EventStream) *job_queue_processor.JobQueueProcessor {
	jobQueueProcessor, err := job_queue_processor.ConstructJobQueueProcessor()
	if err != nil {
		panic(err)
	}

	jobQueueProcessor.Register(eventStream)
	return jobQueueProcessor
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
