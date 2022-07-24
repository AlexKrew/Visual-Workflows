package entities

import (
	"visualWorkflows/internal/storage"
)

type WorkflowID = string

type Workflow struct {
	ID   WorkflowID
	Name string

	Nodes map[NodeID]Node
	Edges map[EdgeID]Edge
}

// WorkflowFromDefinition is a mapper function
// for mapping WorkflowDefinition elements to Workflow elements.
func WorkflowFromDefinition(definition storage.WorkflowDefinition) (Workflow, error) {
	workflow := Workflow{}

	workflow.ID = definition.ID
	workflow.Name = definition.Name
	workflow.Nodes = make(map[string]Node)
	workflow.Edges = make(map[string]Edge)

	for id, node := range definition.Nodes {

		// Mapping of InputPorts
		inputPorts := map[PortID]Port{}
		for id, inputPortDef := range node.InputPorts {
			port, err := InputPortFromDefinition(inputPortDef)
			if err != nil {
				return Workflow{}, err
			}

			inputPorts[id] = port
		}

		// Mapping of OutputPorts
		outputPorts := map[PortID]Port{}
		for id, outputPortDef := range node.OutputPorts {
			port, err := OutputPortFromDefinition(outputPortDef)
			if err != nil {
				return Workflow{}, err
			}

			outputPorts[id] = port
		}

		workflow.Nodes[id] = Node{
			ID:          node.ID,
			Name:        node.Name,
			Type:        node.Type,
			InputPorts:  inputPorts,
			OutputPorts: outputPorts,
		}
	}

	// Mapping of Edges
	for id, edgeDef := range definition.Edges {
		edge, err := EdgeFromDefinition(edgeDef)
		if err != nil {
			return Workflow{}, err
		}

		workflow.Edges[id] = edge
	}

	return workflow, nil
}

func (wf *Workflow) AddNode(node Node) error {
	return nil
}

func (wf *Workflow) AddEdge(origin PortAddress, target PortAddress) error {
	return nil
}
