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
	// workflow.Edges = make(map[string]Edge)

	for _, node := range definition.Nodes {

		ports := []Port{}
		for _, portDef := range node.Ports {
			port, err := PortFromDefinition(portDef)
			if err != nil {
				return Workflow{}, err
			}

			ports = append(ports, port)
		}

		workflow.Nodes = append(workflow.Nodes, Node{
			ID:    node.ID,
			Name:  node.Name,
			Type:  node.Type,
			Ports: ports,
		})
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

func (wf *Workflow) AddNode(node Node) error {
	return nil
}

func (wf *Workflow) AddEdge(origin PortAddress, target PortAddress) error {
	return nil
}
