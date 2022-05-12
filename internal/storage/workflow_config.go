package storage

type WorkflowConfiguration struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Nodes map[string]any `json:"nodes"`
	Edges map[string]any `json:"edges"`
}
