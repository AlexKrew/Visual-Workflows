package workflows

import "workflows/internal/utils"

type NodeID = utils.UUID

type Node struct {
	ID    NodeID `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Ports Ports  `json:"ports"`
}

func (node *Node) PortByID(portId PortID) (Port, bool) {
	for _, port := range node.Ports {
		if port.ID == portId {
			return port, true
		}
	}

	return Port{}, false
}
