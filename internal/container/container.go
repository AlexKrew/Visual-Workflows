package container

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/runtime"
	"visualWorkflows/internal/storage"
	"visualWorkflows/shared/entities"

	wc "visualWorkflows/workerclient"
)

// WorkflowContainer is a fassade between the webserver and the local
// storage and workflow runtimes
type WorkflowContainer struct {
	runtimes map[entities.WorkflowID]*runtime.Runtime
	Events   chan interface{}
}

func Construct() WorkflowContainer {

	runtimes := make(map[entities.WorkflowID]*runtime.Runtime)
	container := WorkflowContainer{
		runtimes: runtimes,
	}

	return container
}

func (container *WorkflowContainer) RegisterWorker(runtimeID string, worker *wc.WorkerClient) {
	r := container.runtimes[runtimeID]
	r.RegisterWorker(worker)
}

/* Use Cases */

func (container *WorkflowContainer) CreateWorkflow(props storage.CreateWorkflowProps) (string, error) {
	return storage.CreateWorkflow(props)
}

func (container *WorkflowContainer) LoadWorkflow(props storage.LoadWorkflowProps) error {

	wfDefinition, err := storage.LoadWorkflowDefinition(props)
	if err != nil {
		return err
	}

	run := runtime.Runtime{}
	err = runtime.Initialize(&run, wfDefinition)
	if err != nil {
		return err
	}

	go container.registerCallbacks(run.Events)

	container.runtimes[run.Workflow.ID] = &run
	fmt.Println("Loaded workflow", run.Workflow.ID)

	return nil
}

func (container *WorkflowContainer) StartWorkflow(id string) error {

	r, exists := container.runtimes[id]
	if !exists {
		return errors.New("workflow not loaded")
	}
	fmt.Println("---- Starting workflow ----")
	r.Start()

	return nil
}

func (container *WorkflowContainer) GetWorkflowById(id string) (entities.Workflow, error) {
	r, exists := container.runtimes[id]
	if !exists {
		return entities.Workflow{}, errors.New("workflow not loaded")
	}

	return r.Workflow, nil
}

func (c *WorkflowContainer) registerCallbacks(ch chan interface{}) {
	event := <-ch
	// TODO: switch
	c.publishLog(event)

}

func (c *WorkflowContainer) publishLog(log interface{}) {
	c.Events <- log
}

/* Storage */

func (container *WorkflowContainer) GetAvailableWorkflows() ([]storage.WorkflowInfo, error) {
	return storage.GetAvailableWorkflows()
}
