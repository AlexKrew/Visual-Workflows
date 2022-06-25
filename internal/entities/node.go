package entities

type NodeID = string

type Node struct {
	ID          NodeID
	Name        string
	Type        string
	Category    string
	InputPorts  map[PortID]Port
	OutputPorts map[PortID]Port
}

func (node *Node) GetInputPortIds() []PortID {
	portIds := []PortID{}

	for id := range node.InputPorts {
		portIds = append(portIds, id)
	}

	return portIds
}

func (node *Node) GetInputPort(portID string) (Port, bool) {
	port, exists := node.InputPorts[portID]
	return port, exists
}

func (node *Node) GetOutputPort(portID string) (Port, bool) {
	port, exists := node.OutputPorts[portID]
	return port, exists
}
