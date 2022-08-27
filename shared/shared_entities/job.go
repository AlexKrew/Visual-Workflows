package shared_entities

import (
	"encoding/json"
	"workflows/internal/utils"
)

type JobID = utils.UUID

type Job struct {
	ID         JobID      `json:"id"`
	NodeID     string     `json:"node_id"`
	Type       string     `json:"type"`
	Input      JobPayload `json:"input"`
	WorkflowID string     `json:"workflow_id"`
	// Wheter the job is open for execution for a worker.
	// Should only be used inside the JobQueue
	Locked bool `json:"locked"`
}

func NewJob(nodeType string, input JobPayload, nodeId string, workflowId string) Job {
	return Job{
		ID:         utils.GetNewUUID(),
		NodeID:     nodeId,
		Type:       nodeType,
		Input:      input,
		WorkflowID: workflowId,
	}
}

func (job *Job) ToJSONString() (string, error) {
	asJson, err := json.Marshal(job)
	if err != nil {
		return "", err
	}

	return string(asJson), nil
}

func JobFromJSONString(jobString string) (Job, error) {
	var job Job
	if err := json.Unmarshal([]byte(jobString), &job); err != nil {
		return job, err
	}

	return job, nil
}
