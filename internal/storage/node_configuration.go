package storage

// NodeConfig describes the default configuration of a usable workflow node
type NodeConfiguration struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
}
