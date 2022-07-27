package entities

import "visualWorkflows/internal/storage"

type EdgeID = string

type Edge struct {
	ID     EdgeID      `json:"id"`
	Origin PortAddress `json:"origin"`
	Target PortAddress `json:"target"`
}

func EdgeFromDefinition(definition storage.EdgeDefinition) (Edge, error) {
	origin := PortAddress{
		NodeID: definition.Origin.NodeID,
		PortID: definition.Origin.PortID,
	}
	target := PortAddress{
		NodeID: definition.Target.NodeID,
		PortID: definition.Target.PortID,
	}

	edge := Edge{
		ID:     definition.ID,
		Origin: origin,
		Target: target,
	}

	return edge, nil
}

func (edge *Edge) ToDefinition() (storage.EdgeDefinition, error) {

	def := storage.EdgeDefinition{
		ID: edge.ID,
		Origin: storage.EdgeEnd{
			NodeID: edge.Origin.NodeID,
			PortID: edge.Origin.PortID,
		},
		Target: storage.EdgeEnd{
			NodeID: edge.Target.NodeID,
			PortID: edge.Target.PortID,
		},
	}

	return def, nil
}
