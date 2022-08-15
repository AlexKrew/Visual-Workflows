package workers

import "workflows/clients/go/pkg/nodes"

func NewInjectWorker() JobWorker {
	return JobWorker{
		JobType: "Inject",
		Handler: nodes.ProcessInject,
	}
}
