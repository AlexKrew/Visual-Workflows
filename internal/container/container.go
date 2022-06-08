package container

import (
	runtime "visualWorkflows/internal/runtime"
	storage "visualWorkflows/internal/storage"
)

// WorkflowContainer is a fassade between the webserver and the local
// storage and workflow runtimes
type WorkflowContainer struct {
	workflows *[]runtime.Runtime
}

func Construct() WorkflowContainer {
	workflows := []runtime.Runtime{}

	container := WorkflowContainer{
		workflows: &workflows,
	}

	return container
}

/* Storage */

func (container *WorkflowContainer) GetAvailableWorkflows() ([]storage.WorkflowInfo, error) {
	return storage.GetAvailableWorkflows()
}
