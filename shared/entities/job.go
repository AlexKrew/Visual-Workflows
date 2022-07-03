package entities

import (
	shared "visualWorkflows/shared/utils"

	"github.com/google/uuid"
)

type JobID = uuid.UUID
type JobType = int16

const (
	JTProcess JobType = iota
)

type Job struct {
	ID      JobID
	Type    JobType
	Payload any
}

type ProcessJobProps struct {
	NodeID   NodeID                     `json:"nodeId"`
	NodeType string                     `json:"nodeType"`
	Input    map[PortID]WorkflowMessage `json:"input"`
}

func BuildProcessJob(payload ProcessJobProps) Job {
	return Job{
		ID:      shared.GetNewUUID(),
		Type:    JTProcess,
		Payload: payload,
	}
}
