package webserver

import (
	"workflows/internal/processors/workflow_processor"
	"workflows/internal/workflows"
)

type WorkflowHelper struct {
	workflowProcessor *workflow_processor.WorkflowProcessor
}

func ConstructWorkflowHelper(processor *workflow_processor.WorkflowProcessor) WorkflowHelper {
	return WorkflowHelper{
		workflowProcessor: processor,
	}
}

func (helper *WorkflowHelper) GetAvailableWorkflows() {
	// TODO:
}

func (helper *WorkflowHelper) WorkflowById(id workflows.WorkflowID) (*workflows.Workflow, bool) {
	workflow, exists := helper.workflowProcessor.WorkflowByID(id)
	if !exists {
		return nil, false
	}
	return workflow, true
}

func (helper *WorkflowHelper) CreateNewWorkflow(name string) {
	command := workflows.CreateWorkflowInstanceCommandBody{
		WorkflowID: name,
	}
	helper.workflowProcessor.EventStream.AddCommand(workflows.NewCreateWorkflowInstanceCommand(command))
}

func (helper *WorkflowHelper) PublishChanges(workflow workflows.Workflow) error {

	err := workflows.WorkflowToFilesystem(workflow)
	if err != nil {
		return err
	}

	container, exists := helper.workflowProcessor.Containers[workflow.ID]
	if !exists {
		panic("workflow does not exist")
	}

	storedWorkflow, err := workflows.WorkflowFromFilesystem(workflow.ID)
	if err != nil {
		return err
	}

	container.Run(&storedWorkflow)

	return nil
}
