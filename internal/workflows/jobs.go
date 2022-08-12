package workflows

import (
	"workflows/internal/utils"
)

type JobID = utils.UUID

type Job struct {
	ID         JobID
	NodeType   string
	NodeID     NodeID
	Input      map[PortID]Message
	Locked     bool
	WorkflowID WorkflowID
}

func NewJob(nodeType string, input map[PortID]Message, nodeId NodeID) Job {
	return Job{
		ID:       utils.GetNewUUID(),
		NodeID:   nodeId,
		NodeType: nodeType,
		Input:    input,
	}
}
