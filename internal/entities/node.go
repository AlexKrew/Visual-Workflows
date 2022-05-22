package entities

import "fmt"

type Node struct {
	ID       string                       `json:"id"`
	Name     string                       `json:"name"`
	Type     string                       `json:"type"`
	Category string                       `json:"category"`
	InPorts  map[PortID]PortConfiguration `json:"inports"`
	OutPorts map[PortID]PortConfiguration `json:"outports"`
}

type PortConfiguration struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

type PortAddress struct {
	NodeID string `json:"node"`
	PortID string `json:"port"`
}
type PortID = string
type UniquePortID = string

func (node *Node) GetInPort(portID string) (PortConfiguration, bool) {
	port, exists := node.InPorts[portID]
	return port, exists
}

func (node *Node) GetOutPort(portID string) (PortConfiguration, bool) {
	port, exists := node.OutPorts[portID]
	return port, exists
}

func (port *PortConfiguration) IsCompatibleWith(other PortConfiguration) bool {
	return port.Type == other.Type
}

func (pa *PortAddress) GetUniquePortID() UniquePortID {
	return fmt.Sprintf("%s-%s", pa.NodeID, pa.PortID)
}
