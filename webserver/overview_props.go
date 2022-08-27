package webserver

import "workflows/internal/workflows"

type CreateWorkflowRequest struct {
	Name     string             `json:"name"`
	Workflow workflows.Workflow `json:"workflow"`
}
