package workflows

import (
	"errors"
	"workflows/internal/utils"
)

type NodeID = utils.UUID

type Node struct {
	ID       NodeID `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	IsUINode bool   `json:"is_ui_node"`
	Ports    Ports  `json:"ports"`
	UI       NodeUI `json:"ui"`
}

func (node *Node) PortByID(portId PortID) (Port, bool) {
	for _, port := range node.Ports {
		if port.ID == portId {
			return port, true
		}
	}

	return Port{}, false
}

func (node *Node) TriggerOutputPortID() (PortID, error) {

	for _, port := range node.Ports {
		if port.IsTrigger && !port.IsInputPort {
			return port.ID, nil
		}
	}

	return "", errors.New("missing trigger port")
}

func (node *Node) PortByIdentifier(identifier string) (PortID, bool) {
	for _, port := range node.Ports {
		if port.Identifier == identifier {
			return port.ID, true
		}
	}

	return "", false
}
