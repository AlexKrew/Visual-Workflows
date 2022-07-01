package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/entities"
)

type MessageRouter struct {
	invertedIndex map[entities.UniquePortID]([]entities.PortAddress)
	store         *MessageStore
}

func constructMessageRouter(workflow entities.Workflow, ms *MessageStore) (*MessageRouter, error) {
	router := MessageRouter{
		invertedIndex: make(map[entities.UniquePortID][]entities.PortAddress),
		store:         ms,
	}

	edges := workflow.Edges
	nodes := workflow.Nodes
	if err := router.buildInvertedIndex(edges, nodes); err != nil {
		return &MessageRouter{}, err
	}

	return &router, nil
}

func (router *MessageRouter) buildInvertedIndex(edges map[string]entities.Edge, nodes map[string]entities.Node) error {

	for _, edge := range edges {

		// Check for existing nodes
		originNode, originNodeExists := nodes[edge.Origin.NodeID]
		targetNode, targetNodeExists := nodes[edge.Target.NodeID]
		if !originNodeExists || !targetNodeExists {
			return errors.New("origin and/or target node of an edge is missing in the workflow")
		}

		// Check for existing ports
		originPort, originPortExists := originNode.GetOutputPort(edge.Origin.PortID)
		targetPort, targetPortExists := targetNode.GetInputPort(edge.Target.PortID)
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

		if router.invertedIndex[origin] == nil {
			router.invertedIndex[origin] = make([]entities.PortAddress, 0)
		}

		router.invertedIndex[origin] = append(router.invertedIndex[origin], target)
		fmt.Println("Added index for", origin)
	}

	return nil
}

// TODO: A publish messages function that takes in a map of OutputPortIds: WFMessage and distributes
// all message to the correct ports in one step

// publishMessage distributes a message to all connected ports defined by origin
func (router *MessageRouter) publishMessage(origin entities.UniquePortID, message entities.WorkflowMessage) error {
	connectedPorts, err := router.getConnectedPorts(origin)
	if err != nil {
		return err
	}

	for _, target := range connectedPorts {
		// TODO: Error handling
		router.store.StoreNewMessage(target, message)
	}

	return nil
}

func (router *MessageRouter) getConnectedPorts(origin entities.UniquePortID) ([]entities.PortAddress, error) {
	ports, exists := router.invertedIndex[origin]
	if !exists {
		return []entities.PortAddress{}, errors.New("Fail")
	}

	return ports, nil
}
