package entities

import (
	"visualWorkflows/internal/storage"
)

type WorkflowID = string

type Workflow struct {
	ID   WorkflowID `json:"id"`
	Name string     `json:"name"`

	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// WorkflowFromDefinition is a mapper function
// for mapping WorkflowDefinition elements to Workflow elements.
func WorkflowFromDefinition(definition storage.WorkflowDefinition) (Workflow, error) {
	workflow := Workflow{}

	workflow.ID = definition.ID
	workflow.Name = definition.Name
	workflow.Nodes = []Node{}
	workflow.Edges = []Edge{}

	for _, nodeDef := range definition.Nodes {

		node, err := NodeFromDefinition(nodeDef)
		if err != nil {
			return Workflow{}, nil
		}

		workflow.Nodes = append(workflow.Nodes, node)
	}

	// Mapping of Edges
	for _, edgeDef := range definition.Edges {
		edge, err := EdgeFromDefinition(edgeDef)
		if err != nil {
			return Workflow{}, err
		}

		workflow.Edges = append(workflow.Edges, edge)
	}

	return workflow, nil
}

func (workflow *Workflow) ToDefinition() (storage.WorkflowDefinition, error) {

	nodes := []storage.NodeDefinition{}
	for _, node := range workflow.Nodes {
		nodeDef, err := node.ToDefinition()
		if err != nil {
			return storage.WorkflowDefinition{}, err
		}

		nodes = append(nodes, nodeDef)
	}

	edges := []storage.EdgeDefinition{}
	for _, edge := range workflow.Edges {
		edgeDef, err := edge.ToDefinition()
		if err != nil {
			return storage.WorkflowDefinition{}, err
		}

		edges = append(edges, edgeDef)
	}

	workflowDef := storage.WorkflowDefinition{
		ID:    workflow.ID,
		Name:  workflow.Name,
		Nodes: nodes,
		Edges: edges,
	}

	return workflowDef, nil
}

func (wf *Workflow) AddNode(node Node) error {
	return nil
}

func (wf *Workflow) AddEdge(origin PortAddress, target PortAddress) error {
	return nil
}
