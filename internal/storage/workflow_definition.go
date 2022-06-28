package storage

type WorkflowDefinition struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Nodes map[string]NodeDefinition
	Edges map[string]EdgeDefinition
}

type NodeDefinition struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	InputPorts  map[string]InputPortDefinition
	OutputPorts map[string]OutputPortDefinition
}

type EdgeDefinition struct {
	ID         string `json:"id"`
	OriginNode string `json:"originNode"`
	OriginPort string `json:"originPort"`
	TargetNode string `json:"targetNode"`
	TargetPort string `json:"targetPort"`
}

type InputPortDefinition struct {
	Label    string `json:"label"`
	DataType string `json:"datatype"`
}

type OutputPortDefinition = InputPortDefinition
