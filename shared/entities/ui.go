package entities

import "visualWorkflows/internal/storage"

type UI struct {
	Position []int `json:"position"`
}

func UIFromDefinition(uiDef storage.UIDefinition) UI {
	return UI{
		Position: uiDef.Position,
	}
}
