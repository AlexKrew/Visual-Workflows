package workflows

import "workflows/internal/utils"

type WorkflowID = utils.UUID

type Workflow struct {
	ID    WorkflowID `json:"id"`
	Name  string     `json:"name"`
	Nodes []Node     `json:"nodes"`
	Edges []Edge     `json:"edges"`
}

func (workflow *Workflow) NodeByID(id NodeID) (Node, bool) {
	for _, node := range workflow.Nodes {
		if node.ID == id {
			return node, true
		}
	}

	return Node{}, false
}
