package workflows

import "workflows/internal/utils"

type WorkflowID = utils.UUID

type Workflow struct {
	ID    WorkflowID `json:"id"`
	Name  string     `json:"name"`
	Nodes []Node     `json:"nodes"`
	Edges []Edge     `json:"edges"`
}

func NewWorkflow(name string) Workflow {
	return Workflow{
		ID:    utils.GetNewUUID(),
		Name:  name,
		Nodes: []Node{},
		Edges: []Edge{},
	}
}

func (workflow *Workflow) NodeByID(id NodeID) (Node, bool) {
	for _, node := range workflow.Nodes {
		if node.ID == id {
			return node, true
		}
	}

	return Node{}, false
}
