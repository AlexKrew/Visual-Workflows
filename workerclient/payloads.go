package workerclient

import "visualWorkflows/shared/entities"

type ProcessNodePayload struct {
	NodeID   entities.NodeID                              `json:"nodeId`
	NodeType string                                       `json:"nodeType"`
	Input    map[entities.PortID]entities.WorkflowMessage `json:"input"`
}
