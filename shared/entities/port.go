package entities

import (
	"fmt"
	"visualWorkflows/internal/storage"
)

type PortID = string
type GroupPortID = string
type UniquePortID = string

type Port struct {
	ID           PortID
	Label        string
	DataType     string
	DefaultValue WorkflowMessage
	Added        bool
	IsInput      bool
}

type PortAddress struct {
	NodeID string
	PortID string
}

func PortFromDefinition(definition storage.PortDefinition) (Port, error) {

	port := Port{
		ID:           definition.ID,
		Label:        definition.Label,
		DataType:     definition.DataType,
		DefaultValue: EmptyMessage(),
		Added:        definition.Added,
		IsInput:      definition.IsInput,
	}

	return port, nil
}

func (port *Port) IsCompatibleWith(other Port) bool {
	return port.DataType == other.DataType
}

func (port *Port) GetMessage() WorkflowMessage {
	return port.DefaultValue
}

func (pa *PortAddress) GetUniquePortID() UniquePortID {
	return fmt.Sprintf("%s-%s", pa.NodeID, pa.PortID)
}
