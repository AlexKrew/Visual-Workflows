package storage

type WorkflowDefinition struct {
	ID    string           `json:"id"`
	Name  string           `json:"name"`
	Nodes []NodeDefinition `json:"nodes"`
	Edges []EdgeDefinition `json:"edges"`
}

type NodeDefinition struct {
	ID    string           `json:"id"`
	Name  string           `json:"name"`
	Type  string           `json:"type"`
	Ports []PortDefinition `json:"ports"`
	UI    UIDefinition     `json:"ui"`
}

type PortDefinition struct {
	ID           string `json:"id"`
	Label        string `json:"label"`
	DataType     string `json:"datatype"`
	IsInput      bool   `json:"is_input"`
	Added        bool   `json:"added"`
	DefaultValue any    `json:"default_value"`
}

type UIDefinition struct {
	Position []int `json:"position"`
}

type EdgeDefinition struct {
	ID     string  `json:"id"`
	Origin EdgeEnd `json:"origin"`
	Target EdgeEnd `json:"target"`
}

type EdgeEnd struct {
	NodeID string `json:"node_id"`
	PortID string `json:"port_id"`
}

// type InputPortDefinition struct {
// 	Label    string `json:"label"`
// 	DataType string `json:"datatype"`
// 	// TODO: Add message field
// }
// type OutputPortDefinition struct {
// 	Label    string `json:"label"`
// 	DataType string `json:"datatype"`
// 	// TODO: Add message field
// }

// type OutputPortDefinition = InputPortDefinition
