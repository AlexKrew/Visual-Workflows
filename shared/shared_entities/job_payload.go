package shared_entities

type JobPayloadItem struct {
	NodeID         string          `json:"node_id"`
	PortIdentifier string          `json:"port_identifier"`
	GroupID        string          `json:"group_id"` // Empty string if port does not belong to a group
	Value          WorkflowMessage `json:"value"`
}

type JobPayload = []JobPayloadItem
