package workflows

import (
	"time"
	"workflows/internal/utils"
)

type WorkflowEventType = int

type WorkflowEvent struct {
	ID        utils.UUID        `json:"id"`
	Type      WorkflowEventType `json:"type"`
	Body      any               `json:"body"`
	CreatedAt int64             `json:"created_at"`
}

const (
	WorkflowInstanceCreated = iota
	WorkflowReady
	WorkflowStarted
	WorkflowStopped

	// Jobs
	JobCreated
	JobCompleted

	// Debug
	DebugEvent
)

func createWorkflowEvent(event WorkflowEventType, body any) WorkflowEvent {
	return WorkflowEvent{
		ID:        utils.GetNewUUID(),
		Type:      event,
		Body:      body,
		CreatedAt: time.Now().UnixMilli(),
	}
}

// ---- Workflow Instance Created ----
type WorkflowInstanceCreatedEventBody struct {
	Workflow Workflow
}

func NewWorkflowInstanceCreatedEvent(body WorkflowInstanceCreatedEventBody) WorkflowEvent {
	return createWorkflowEvent(WorkflowInstanceCreated, body)
}

// ---- Workflow Ready ----
type WorkflowReadyEventBody struct {
	WorkflowID WorkflowID
}

func NewWorkflowReadyEvent(body WorkflowReadyEventBody) WorkflowEvent {
	return createWorkflowEvent(WorkflowReady, body)
}

// ---- Workflow Started ----
type WorkflowStartedEventBody struct {
	WorkflowID WorkflowID
}

func NewWorkflowStartedEvent(body WorkflowStartedEventBody) WorkflowEvent {
	return createWorkflowEvent(WorkflowStarted, body)
}

// ---- Workflow Stopped ----
type WorkflowStoppedEventBody struct {
	WorkflowID WorkflowID
}

func NewWorkflowStoppedEvent(body WorkflowStartedEventBody) WorkflowEvent {
	return createWorkflowEvent(WorkflowStopped, body)
}

// ---- Job Created ----
type JobCreatedEventBody struct {
	WorkflowID WorkflowID
	Job        Job
}

func NewJobCreatedEvent(body JobCreatedEventBody) WorkflowEvent {
	return createWorkflowEvent(JobCreated, body)
}

// ---- Job Completed ----
type JobCompletedEventBody struct {
	WorkflowID WorkflowID
	Job        Job
	Result     any
}

func NewJobCompletedEvent(body JobCompletedEventBody) WorkflowEvent {
	return createWorkflowEvent(JobCompleted, body)
}

// ---- Debug ----
type DebugEventBody struct {
	WorkflowID WorkflowID
	Value      any
}

func NewDebugEvent(body DebugEventBody) WorkflowEvent {
	return createWorkflowEvent(DebugEvent, body)
}
