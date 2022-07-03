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
		NodeID: definition.OriginNode,
		PortID: definition.OriginPort,
	}
	target := PortAddress{
		NodeID: definition.TargetNode,
		PortID: definition.TargetPort,
	}

	edge := Edge{
		ID:     definition.ID,
		Origin: origin,
		Target: target,
	}

	return edge, nil
}
