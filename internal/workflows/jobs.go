package workflows

import (
	"workflows/internal/utils"
)

type JobID = utils.UUID

type Job struct {
	ID       JobID
	NodeType string
	Input    map[PortID]Message
	Locked   bool
}

func NewJob(nodeType string, input map[PortID]Message) Job {
	return Job{
		ID:       utils.GetNewUUID(),
		NodeType: nodeType,
		Input:    input,
	}
}
