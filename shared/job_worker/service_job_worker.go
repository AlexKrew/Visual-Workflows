package job_worker

import (
	"log"
	"workflows/shared/nodes"
	"workflows/shared/shared_entities"
)

type ServiceHandlerFunc func(input *nodes.NodeInput, output *nodes.NodeOutput) error

// ServiceJobWorker is an implementation of JobWorker
type ServiceJobWorker struct {
	Type    string
	Handler ServiceHandlerFunc
}

func NewServiceJobWorker(jobType string, handler ServiceHandlerFunc) ServiceJobWorker {
	return ServiceJobWorker{
		Type:    jobType,
		Handler: handler,
	}
}

func (worker ServiceJobWorker) JobType() string {
	return worker.Type
}

func (worker ServiceJobWorker) ProcessJob(job shared_entities.Job) shared_entities.JobResult {

	var result shared_entities.JobResult

	input := nodes.NodeInputFromJob(job)
	output := nodes.NewNodeOutput()

	err := worker.Handler(&input, &output)
	if err != nil {
		log.Printf("failed to handle job: %s", err.Error())
		return result
	}

	result = shared_entities.JobResult{
		JobID:  job.ID,
		NodeID: job.NodeID,
		Output: output.GetOutput(),
		Logs:   output.GetLogs(),
	}

	return result
}
