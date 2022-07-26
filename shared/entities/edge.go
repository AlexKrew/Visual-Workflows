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
		NodeID: definition.Origin.Node,
		PortID: definition.Origin.Port,
	}
	target := PortAddress{
		NodeID: definition.Target.Node,
		PortID: definition.Target.Port,
	}

	edge := Edge{
		ID:     definition.ID,
		Origin: origin,
		Target: target,
	}

	return edge, nil
}
