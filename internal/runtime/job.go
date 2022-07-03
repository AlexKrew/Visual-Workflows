package runtime

import (
	"visualWorkflows/internal/entities"

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
	NodeID entities.NodeID `json:"id"`
	Input  map[entities.PortID]entities.WorkflowMessage
}

func buildProcessJob(payload ProcessJobProps) Job {
	return Job{
		ID:      getNewUUID(),
		Type:    JTProcess,
		Payload: payload,
	}
}
