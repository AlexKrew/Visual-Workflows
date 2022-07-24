package runtime

import (
	"visualWorkflows/internal/storage"
	"visualWorkflows/shared/entities"
	wc "visualWorkflows/workerclient"
)

type Runtime struct {
	initialized bool

	Workflow      entities.Workflow
	Router        *MessageRouter
	Store         *MessageStore
	EventStreamer *EventStreamer
	JobQueue      *JobQueue

	// TODO: Logger

	// Only Runtime Events - not the same as the events in EventStreamer
	Events chan interface{}

	// TODO: Move workers map to container
	Workers []*wc.WorkerClient
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
	runtime.Workers = make([]*wc.WorkerClient, 0)
	runtime.Events = make(chan interface{})

	runtime.initialized = true

	registerOperations(eventStreamer)

	return nil
}

func (r *Runtime) RegisterWorker(worker *wc.WorkerClient) {
	r.Workers = append(r.Workers, worker)
}

func registerOperations(eventStreamer *EventStreamer) {
	go eventStreamer.addOperation(createJobOperation)
	go eventStreamer.addOperation(handleJobResultOperation)
}

func (r *Runtime) Start() {
	r.Store.StoreNewMessage(entities.PortAddress{
		NodeID: "node1",
		PortID: "input",
	}, entities.BooleanMessage(true))
}

func (r *Runtime) Stop() {

}
