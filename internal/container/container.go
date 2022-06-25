package container

import (
	"fmt"
	"visualWorkflows/internal/entities"
	"visualWorkflows/internal/runtime"
	"visualWorkflows/internal/storage"
)

// WorkflowContainer is a fassade between the webserver and the local
// storage and workflow runtimes
type WorkflowContainer struct {
	workflows map[entities.WorkflowID]runtime.Runtime
}

func Construct() WorkflowContainer {
	workflows := map[entities.WorkflowID]runtime.Runtime{}

	container := WorkflowContainer{
		workflows: workflows,
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

	runtime, err := runtime.Initialize(wfDefinition)
	if err != nil {
		return err
	}

	container.workflows[runtime.Workflow.ID] = runtime
	fmt.Println("Loaded workflow", runtime.Workflow.ID)

	return nil
}

/* Storage */

func (container *WorkflowContainer) GetAvailableWorkflows() ([]storage.WorkflowInfo, error) {
	return storage.GetAvailableWorkflows()
}
