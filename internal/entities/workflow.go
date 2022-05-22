package entities

type Workflow struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Nodes map[string]Node `json:"nodes"`
	Edges map[string]Edge `json:"edges"`
}
