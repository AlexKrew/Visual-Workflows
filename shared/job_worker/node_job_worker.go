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
		NodeID:      result.NodeID,
		Output:      outputValues,
		Logs:        output.GetLogs(),
		DontTrigger: output.DontTrigger(),
	}

	return result
}
