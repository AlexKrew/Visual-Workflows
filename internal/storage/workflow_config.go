package storage

type WorkflowConfiguration struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Nodes map[string]nodeConfig `json:"nodes"`
	Edges map[string]any        `json:"edges"`
}

type nodeConfig = map[string]any
