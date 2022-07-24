package entities

import "visualWorkflows/shared/utils"

type JobResult struct {
	ID     utils.UUID
	Logs   []string
	Output map[PortID]WorkflowMessage
	Errors []error
}

func JobResultError(id utils.UUID, err error) JobResult {
	return JobResult{
		ID:     id,
		Errors: []error{err},
	}
}
