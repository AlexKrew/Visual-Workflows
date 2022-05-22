package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/entities"
)

type MessageRouter struct {
	InvertedIndex map[entities.UniquePortID]([]entities.PortAddress)
}

func constructMessageRouter(rt *Runtime) (MessageRouter, error) {
	router := MessageRouter{}
	router.InvertedIndex = make(map[entities.UniquePortID][]entities.PortAddress)

	edges := rt.Workflow.Edges
	nodes := rt.Workflow.Nodes

	err := router.generateInvertedIndex(edges, nodes)
	if err != nil {
		return MessageRouter{}, err
	}

	return router, nil
}

func (router *MessageRouter) generateInvertedIndex(edges map[string]entities.Edge, nodes map[string]entities.Node) error {

	for _, edge := range edges {
		fmt.Println("Edge:", edge.ID)

		// Check for existing nodes
		originNode, originNodeExists := nodes[edge.Origin.NodeID]
		targetNode, targetNodeExists := nodes[edge.Target.NodeID]
		if !originNodeExists || !targetNodeExists {
			return errors.New("origin and/or target node of an edge is missing in the workflow")
		}

		// Check for existing ports
		originPort, originPortExists := originNode.GetOutPort(edge.Origin.PortID)
		targetPort, targetPortExists := targetNode.GetInPort(edge.Target.PortID)
		if !originPortExists || !targetPortExists {
			return errors.New("origin and/or target port of an edge does not exist")
		}

		// Check if ports are compatible
		portsCompatible := originPort.IsCompatibleWith(targetPort)
		if !portsCompatible {
			return errors.New("ports are not compatible")
		}

		// Edge is valid
		// Add connection to index
		origin := edge.Origin.GetUniquePortID()
		target := edge.Target

		if router.InvertedIndex[origin] == nil {
			router.InvertedIndex[origin] = make([]entities.PortAddress, 0)
		}

		router.InvertedIndex[origin] = append(router.InvertedIndex[origin], target)
		fmt.Println("Added index for", origin)
	}

	return nil
}

// publishMessage sends messages from an OutPort of a node to all connected nodes
// func (router *MessageRouter) publishMessage(origin UniquePortID) error {
// 	return nil
// }
