package entities

import (
	"fmt"
	"visualWorkflows/internal/storage"
)

type NodeID = string

type Node struct {
	ID    NodeID `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Ports []Port `json:"ports"`
	UI    UI     `json:"ui"`
}

func NodeFromDefinition(definition storage.NodeDefinition) (Node, error) {

	ports := []Port{}
	for _, portDef := range definition.Ports {
		port, err := PortFromDefinition(portDef)
		if err != nil {
			return Node{}, err
		}

		ports = append(ports, port)
	}

	node := Node{
		ID:    definition.ID,
		Name:  definition.Name,
		Type:  definition.Type,
		Ports: ports,
		UI:    UIFromDefinition(definition.UI),
	}

	return node, nil
}

func (node *Node) ToDefinition() (storage.NodeDefinition, error) {

	ports := []storage.PortDefinition{}
	for _, port := range node.Ports {
		portDef, err := port.ToDefinition()
		if err != nil {
			return storage.NodeDefinition{}, err
		}

		ports = append(ports, portDef)
	}

	def := storage.NodeDefinition{
		ID:    node.ID,
		Name:  node.Name,
		Type:  node.Type,
		Ports: ports,
		UI:    storage.UIDefinition(node.UI),
	}

	return def, nil
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
