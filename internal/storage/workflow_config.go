package storage

type WorkflowConfiguration struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Nodes map[string]NodeDefinition `json:"nodes"`
	Edges map[string]any            `json:"edges"`
}

type NodeDefinition = map[string]any
