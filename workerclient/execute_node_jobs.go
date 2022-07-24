package workerclient

import (
	"errors"
	"fmt"
	"visualWorkflows/nodes"
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/node"
)

func executeNodeJobs(job entities.Job) (entities.JobResult, error) {

	if job.Type == entities.JTProcess {
		return processNode(job), nil
	}

	return entities.JobResult{}, errors.New("unknown node job type")
}

func processNode(job entities.Job) entities.JobResult {

	// TODO: Conversion from json
	payload := job.Payload.(entities.ProcessJobProps)

	nodeType := payload.NodeType
	messages := payload.Input

	fmt.Println("Execute Node", nodeType)

	input := entities.Input{
		Messages: messages,
	}
	output := &entities.Output{
		Messages: make(map[string]entities.WorkflowMessage),
	}
	ctxProxy := node.ConstructWorkflowContext()

	// TODO: move somewhere else
	switch nodeType {
	case "debug":
		nodes.ProcessDebug(input, output, &ctxProxy)
	case "http-request":
		nodes.ProcessHttpRequest(input, output, &ctxProxy)
	}

	result := entities.JobResult{
		ID:     job.ID,
		NodeId: job.NodeId,
		Logs:   ctxProxy.GetLogs(),
		Output: output.Messages,
		Errors: []error{},
	}

	return result
}
