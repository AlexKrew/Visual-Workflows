package entities

import "visualWorkflows/shared/utils"

type JobResult struct {
	ID     utils.UUID
	NodeId NodeID
	Logs   []string
	Output map[PortID]WorkflowMessage
	Errors []error
}

func JobResultError(origin Job, err error) JobResult {
	return JobResult{
		ID:     origin.ID,
		NodeId: origin.NodeId,
		Errors: []error{err},
	}
}
