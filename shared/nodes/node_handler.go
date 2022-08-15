package nodes

import (
	"errors"
	"workflows/shared/shared_entities"
)

type NodeInput struct {
	input map[string]shared_entities.WorkflowMessage
}

func NodeInputFromJob(job shared_entities.Job) NodeInput {
	return NodeInput{
		input: job.Input,
	}
}

func (in *NodeInput) ValueFor(key string) (*shared_entities.WorkflowMessage, error) {
	message, ok := in.input[key]
	if !ok {
		return nil, errors.New("missing key")
	}

	return &message, nil
}

type NodeOutput struct {
	output map[string]shared_entities.WorkflowMessage
	log    []string
}

func NewNodeOutput() NodeOutput {
	return NodeOutput{
		output: make(map[string]shared_entities.WorkflowMessage),
		log:    []string{},
	}
}

func (out *NodeOutput) Set(key string, value shared_entities.WorkflowMessage) {
	out.output[key] = value
}

func (out *NodeOutput) Log(msg string) {
	out.log = append(out.log, msg)
}

func (out *NodeOutput) GetOutput() map[string]shared_entities.WorkflowMessage {
	return out.output
}

func (out *NodeOutput) GetLogs() []string {
	return out.log
}
