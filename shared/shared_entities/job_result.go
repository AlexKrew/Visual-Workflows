package shared_entities

import (
	"encoding/json"
)

type JobResult struct {
	JobID  JobID                      `json:"job_id"`
	NodeID string                     `json:"node_id"`
	Output map[string]WorkflowMessage `json:"output"`
	Logs   []string                   `json:"logs"`
}

func (jobResult *JobResult) ToJSONString() (string, error) {
	asJson, err := json.Marshal(jobResult)
	if err != nil {
		return "", err
	}

	return string(asJson), nil
}

func JobResultFromJSONString(jobResult string) (JobResult, error) {
	var result JobResult

	if err := json.Unmarshal([]byte(jobResult), &result); err != nil {
		return result, err
	}

	return result, nil
}
