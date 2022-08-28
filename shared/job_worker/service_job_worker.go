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

	outputValues := shared_entities.JobPayload{}

	for key, value := range output.GetOutput() {
		outputValues = append(outputValues, shared_entities.JobPayloadItem{
			NodeID:         job.NodeID,
			GroupID:        "",
			PortIdentifier: key,
			Value:          value,
		})
	}

	for groupId, items := range output.GetGroupOutput() {
		for key, value := range items {
			outputValues = append(outputValues, shared_entities.JobPayloadItem{
				NodeID:         job.NodeID,
				GroupID:        groupId,
				PortIdentifier: key,
				Value:          value,
			})
		}
	}

	result = shared_entities.JobResult{
		JobID:       job.ID,
		NodeID:      job.NodeID,
		Output:      outputValues,
		Logs:        output.GetLogs(),
		DontTrigger: output.DontTrigger(),
	}

	return result
}
