package runtime

import (
	"fmt"
	"visualWorkflows/internal/entities"
	"visualWorkflows/internal/storage"
)

type Runtime struct {
	Workflow entities.Workflow
	router   MessageRouter
	store    MessageStore
}

func Initialize(workflowDefinition storage.WorkflowDefinition) (Runtime, error) {

	workflow, err := entities.WorkflowFromDefinition(workflowDefinition)
	if err != nil {
		return Runtime{}, err
	}
	fmt.Println("Constructed workflow from definition")

	store, err := constructMessageStore(workflow)
	if err != nil {
		return Runtime{}, err
	}
	fmt.Println("Constructed message store")

	router, err := constructMessageRouter(workflow, &store)
	if err != nil {
		return Runtime{}, err
	}
	fmt.Println("Constructed message router")

	runtime := Runtime{
		Workflow: workflow,
		router:   router,
		store:    store,
	}

	return runtime, nil
}

func (r *Runtime) Start() {

}

func (r *Runtime) Stop() {

}
