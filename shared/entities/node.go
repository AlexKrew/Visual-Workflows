package entities

import "fmt"

type NodeID = string

type Node struct {
	ID    NodeID `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Ports []Port `json:"ports"`
	UI    UI     `json:"ui"`
}

func (node *Node) GetInputPortIds() []PortID {
	portIds := []PortID{}

	for _, port := range node.Ports {
		portIds = append(portIds, port.ID)
	}

	return portIds
}

func (node *Node) GetInputPort(portID PortID) (Port, bool) {
	port, exists := node.GetPortById(portID)

	fmt.Println("Get input", portID, exists)

	if !exists || !port.IsInput {
		return Port{}, false
	}

	return port, true
}

func (node *Node) GetOutputPort(portID PortID) (Port, bool) {
	port, exists := node.GetPortById(portID)

	fmt.Println("Get output", portID, exists)

	if !exists || port.IsInput {
		return Port{}, false
	}

	return port, true
}

func GetNodeById(nodeID NodeID, nodes []Node) (Node, bool) {
	for _, node := range nodes {
		if node.ID == nodeID {
			return node, true
		}
	}

	return Node{}, false
}

func (node *Node) GetPortById(portID PortID) (Port, bool) {
	for _, port := range node.Ports {
		if port.ID == portID {
			return port, true
		}
	}

	return Port{}, false
}
