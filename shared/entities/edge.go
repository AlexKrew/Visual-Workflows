package entities

import "visualWorkflows/internal/storage"

type EdgeID = string

type Edge struct {
	ID     EdgeID
	Origin PortAddress
	Target PortAddress
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
