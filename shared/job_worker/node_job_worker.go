package job_worker

import (
	"log"
	"workflows/shared/nodes"
	"workflows/shared/shared_entities"
)

type NodeHandlerFunc func(in *nodes.NodeInput, out *nodes.NodeOutput) error

// NodeJobWorker is implementation of JobWorker
type NodeJobWorker struct {
	Type    string
	Handler NodeHandlerFunc
}

func NewNodeJobWorker(jobType string, handler NodeHandlerFunc) NodeJobWorker {
	return NodeJobWorker{
		Type:    jobType,
		Handler: handler,
	}
}

func (worker NodeJobWorker) JobType() string {
	return worker.Type
}

func (worker NodeJobWorker) ProcessJob(job shared_entities.Job) shared_entities.JobResult {

	var result shared_entities.JobResult

	input := nodes.NodeInputFromJob(job)
	output := nodes.NewNodeOutput()

	err := worker.Handler(&input, &output)
	if err != nil {
		log.Printf("failed to handle job: %s", err.Error())
		return result
	}

	result = shared_entities.JobResult{
		ID:     job.ID,
		Output: output.GetOutput(),
		Logs:   output.GetLogs(),
	}

	return result
}
