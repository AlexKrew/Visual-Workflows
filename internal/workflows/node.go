package workflows

import (
	"errors"
	"workflows/internal/utils"
)

type NodeID = utils.UUID

type Node struct {
	ID    NodeID `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Ports Ports  `json:"ports"`
	UI    NodeUI `json:"ui"`
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
