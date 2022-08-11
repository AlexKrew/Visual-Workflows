package workers

import "workflows/clients/go/pkg/nodes"

type JobWorker struct {
	JobType string
	Handler nodes.HandlerFunc
}
