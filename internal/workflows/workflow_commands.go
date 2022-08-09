package workflows

import (
	"time"
	"workflows/internal/utils"
)

type WorkflowCommandType = int

type WorkflowCommand struct {
	ID        utils.UUID          `json:"id"`
	Type      WorkflowCommandType `json:"type"`
	Body      any                 `json:"body"`
	CreatedAt time.Time           `json:"created_at"`
}

const (
	CreateWorkflowInstance = iota
	StartWorkflow
	StopWorkflow

	// Jobs
	CreateJob
	// LockJob
	CompleteJob
)

func createWorkflowCommand(command WorkflowCommandType, body any) WorkflowCommand {
	return WorkflowCommand{
		ID:        utils.GetNewUUID(),
		Type:      command,
		Body:      body,
		CreatedAt: time.Now(),
	}
}

// ---- Create Workflow Instance ----
type CreateWorkflowInstanceCommandBody struct {
	WorkflowID WorkflowID
}

func NewCreateWorkflowInstanceCommand(body CreateWorkflowInstanceCommandBody) WorkflowCommand {
	return createWorkflowCommand(CreateWorkflowInstance, body)
}

// ---- Start Workflow ----
type StartWorkflowCommandBody struct {
	WorkflowID WorkflowID
}

func NewStartWorkflowCommand(body StartWorkflowCommandBody) WorkflowCommand {
	return createWorkflowCommand(StartWorkflow, body)
}

// ---- Stop Workflow ----
type StopWorkflowCommandBody struct {
	WorkflowID WorkflowID
}

func NewStopWorkflowCommand(body StopWorkflowCommandBody) WorkflowCommand {
	return createWorkflowCommand(StopWorkflow, body)
}

// ---- Create Job ----
type CreateJobCommandBody struct {
	WorkflowInstanceID WorkflowContainerID
	NodeID             NodeID
}

func NewCreateJobCommand(body CreateJobCommandBody) WorkflowCommand {
	return createWorkflowCommand(CreateJob, body)
}

// ---- Lock Job ----
// type LockJobCommandBody struct {
// 	WorkflowInstanceID WorkflowContainerID
// 	JobID              JobID
// }

// func NewLockJobCommand(body LockJobCommandBody) WorkflowCommand {
// 	return createWorkflowCommand(LockJob, body)
// }

// ---- Complete Job ----
type CompleteJobCommandBody struct {
	WorflowInstanceID WorkflowContainerID
	JobID             JobID

	// TODO: Check if needed here
	Output any
}

func NewCompletejobCommandBody(body CompleteJobCommandBody) WorkflowCommand {
	return createWorkflowCommand(CompleteJob, body)
}
