package webserver

import (
	"log"
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

func (helper *WorkflowHelper) StartWorkflow(workflowId workflows.WorkflowID) error {
	return helper.workflowProcessor.StartWorkflow(workflowId)
}

func (helper *WorkflowHelper) PublishChanges(workflow workflows.Workflow) error {

	err := workflows.WorkflowToFilesystem(workflow)
	if err != nil {
		return err
	}

	helper.workflowProcessor.ReloadContainer(workflow.ID)
	// container, exists := helper.workflowProcessor.Containers[workflow.ID]
	// if !exists {
	// 	panic("workflow does not exist")
	// }

	// storedWorkflow, err := workflows.WorkflowFromFilesystem(workflow.ID)
	// if err != nil {
	// 	return err
	// }

	// container.Run(&storedWorkflow)

	return nil
}

func (helper *WorkflowHelper) LoadWorkflowById(workflowId workflows.WorkflowID) (*workflows.Workflow, bool) {
	err := helper.workflowProcessor.CreateContainer(workflowId)
	if err != nil {
		log.Panicf("failed to load workflow: %s", err.Error())
		return nil, false
	}

	workflow, exists := helper.WorkflowById(workflowId)
	if !exists {
		log.Panicln("missing workflow. this should never happen")
		return nil, false
	}

	return workflow, true
}
