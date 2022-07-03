package runtime

import (
	"visualWorkflows/internal/entities"
	"visualWorkflows/internal/storage"
)

type Runtime struct {
	initialized bool

	Workflow      entities.Workflow
	Router        *MessageRouter
	Store         *MessageStore
	EventStreamer *EventStreamer
	JobQueue      *JobQueue
}

func Initialize(runtime *Runtime, workflowDefinition storage.WorkflowDefinition) error {

	runtime.initialized = false

	workflow, err := entities.WorkflowFromDefinition(workflowDefinition)
	if err != nil {
		return err
	}
	// fmt.Println("Constructed workflow from definition")

	store, err := constructMessageStore(workflow, runtime)
	if err != nil {
		return err
	}
	// fmt.Println("Constructed message store")

	router, err := constructMessageRouter(workflow, store)
	if err != nil {
		return err
	}
	// fmt.Println("Constructed message router")

	jobQueue, err := constructJobQueue(runtime)
	if err != nil {
		return err
	}

	eventStreamer, err := constructEventStream(runtime)
	if err != nil {
		return err
	}

	runtime.Workflow = workflow
	runtime.Router = router
	runtime.Store = store
	runtime.EventStreamer = eventStreamer
	runtime.JobQueue = jobQueue

	runtime.initialized = true

	registerOperations(eventStreamer)

	return nil
}

func registerOperations(eventStreamer *EventStreamer) {
	go eventStreamer.addOperation(createJobOperation)
}

func (r *Runtime) Start() {
	r.Store.StoreNewMessage(entities.PortAddress{
		NodeID: "node1",
		PortID: "input",
	}, entities.BooleanMessage(true))
}

func (r *Runtime) Stop() {

}
