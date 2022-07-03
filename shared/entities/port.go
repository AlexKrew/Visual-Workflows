package entities

import (
	"fmt"
	"visualWorkflows/internal/storage"
)

type PortID = string
type GroupPortID = string
type UniquePortID = string

type Port struct {
	Label    string
	DataType string
	Message  WorkflowMessage
}

type PortAddress struct {
	NodeID string
	PortID string
}

func PortFromDefinition(definition storage.InputPortDefinition) (Port, error) {

	port := Port{
		Label:    definition.Label,
		DataType: definition.DataType,
		Message:  EmptyMessage(),
	}

	return port, nil
}

func (port *Port) IsCompatibleWith(other Port) bool {
	return port.DataType == other.DataType
}

func (port *Port) GetMessage() WorkflowMessage {
	return port.Message
}

func (pa *PortAddress) GetUniquePortID() UniquePortID {
	return fmt.Sprintf("%s-%s", pa.NodeID, pa.PortID)
}
