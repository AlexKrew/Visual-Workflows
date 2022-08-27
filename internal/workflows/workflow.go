package workflows

import "workflows/internal/utils"

type WorkflowID = utils.UUID

type Workflow struct {
	ID    WorkflowID `json:"id"`
	Name  string     `json:"name"`
	Nodes []Node     `json:"nodes"`
	Edges []Edge     `json:"edges"`
}

func NewWorkflow(name string, workflow Workflow) Workflow {
	newWorkflow := Workflow{
		ID:    utils.GetNewUUID(),
		Name:  name,
		Nodes: workflow.Nodes,
		Edges: workflow.Edges,
	}
	return newWorkflow
}

func (workflow *Workflow) NodeByID(id NodeID) (Node, bool) {
	for _, node := range workflow.Nodes {
		if node.ID == id {
			return node, true
		}
	}

	return Node{}, false
}
