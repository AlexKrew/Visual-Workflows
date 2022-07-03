package node

import (
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/utils"
)

//TODO: Replace logs with shared/entities/logger (if possible)

type WorkflowContextProxy struct {
	logs []string
}

func ConstructWorkflowContext() WorkflowContextProxy {
	return WorkflowContextProxy{
		logs: make([]string, 0),
	}
}

func (wcp *WorkflowContextProxy) Log(message string) {
	wcp.logs = append(wcp.logs, message)
}

func (wcp *WorkflowContextProxy) GetLogs() []string {
	return wcp.logs
}

type ProcessNodeBody struct {
	NodeID   utils.UUID
	NodeType string
	Input    map[entities.PortID]entities.WorkflowMessage
}
