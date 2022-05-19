package entities

type Edge struct {
	ID     string      `json:"id"`
	Origin PortAddress `json:"origin"`
	Target PortAddress `json:"target"`
}
