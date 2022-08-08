package workflows

import "errors"

type MessageRouter struct {
	workflow       *Workflow
	connectedPorts map[UniquePortID][]PortAddress
}

func ConstructMessageRouter(workflow *Workflow) (MessageRouter, error) {

	connectedPorts := make(map[UniquePortID][]PortAddress)

	for _, edge := range workflow.Edges {

		originNode, originNodeExists := workflow.NodeByID(edge.Origin.NodeID)
		targetNode, targetNodeExists := workflow.NodeByID(edge.Target.NodeID)
		if !originNodeExists || !targetNodeExists {
			return MessageRouter{}, errors.New("missing origin/target node")
		}

		originPort, originPortExists := OutputPortByID(edge.Origin.PortID, originNode.Ports)
		targetPort, targetPortExists := InputPortByID(edge.Target.PortID, targetNode.Ports)
		if !originPortExists || !targetPortExists {
			return MessageRouter{}, errors.New("missing origin/target port")
		}

		compatible, compatibilityErrors := originPort.IsCompatibleWith(targetPort)
		if !compatible {
			return MessageRouter{}, compatibilityErrors[0]
		}

		originPortAddr := PortAddress{NodeID: originNode.ID, PortID: originPort.ID}
		uniquePortID := originPortAddr.UniquePortID()
		targetPortAddr := PortAddress{NodeID: targetNode.ID, PortID: targetPort.ID}

		if _, ok := connectedPorts[uniquePortID]; !ok {
			connectedPorts[uniquePortID] = []PortAddress{}
		}
		connectedPorts[uniquePortID] = append(connectedPorts[uniquePortID], targetPortAddr)
	}

	router := MessageRouter{
		workflow:       workflow,
		connectedPorts: connectedPorts,
	}

	return router, nil
}
