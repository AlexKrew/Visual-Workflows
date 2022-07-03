package container

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/entities"
	"visualWorkflows/internal/runtime"
	"visualWorkflows/internal/storage"
)

// WorkflowContainer is a fassade between the webserver and the local
// storage and workflow runtimes
type WorkflowContainer struct {
	runtimes map[entities.WorkflowID]*runtime.Runtime
}

func Construct() WorkflowContainer {

	runtimes := make(map[entities.WorkflowID]*runtime.Runtime)
	container := WorkflowContainer{
		runtimes: runtimes,
	}

	return container
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

	container.runtimes[run.Workflow.ID] = &run
	fmt.Println("Loaded workflow", run.Workflow.ID)

	return nil
}

func (container *WorkflowContainer) StartWorkflow(id string) error {

	r, exists := container.runtimes[id]
	if !exists {
		return errors.New("workflow not loaded")
	}
	r.Start()

	return nil
}

/* Storage */

func (container *WorkflowContainer) GetAvailableWorkflows() ([]storage.WorkflowInfo, error) {
	return storage.GetAvailableWorkflows()
}
