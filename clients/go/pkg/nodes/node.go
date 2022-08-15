package nodes

import (
	"encoding/json"
	"workflows/clients/go/pkg/entities"
)

type HandlerFunc func(input *NodeInput, output *NodeOutput)

type NodeInput struct {
	values map[string]entities.Message
}

func NodeInputFromJob(job entities.ActivatedJob) NodeInput {

	input := NodeInput{
		values: make(map[string]entities.Message),
	}

	json.Unmarshal([]byte(job.Job.Input), &input.values)

	return input
}

type NodeOutput struct {
	outputValues map[string]any
	logs         []string
}

func NewNodeOutput() NodeOutput {
	return NodeOutput{
		outputValues: make(map[string]any),
		logs:         []string{},
	}
}
