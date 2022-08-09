package workflows

import "workflows/internal/utils"

type EdgeID = utils.UUID

type Edge struct {
	ID     EdgeID      `json:"id"`
	Origin PortAddress `json:"origin"`
	Target PortAddress `json:"target"`
}
