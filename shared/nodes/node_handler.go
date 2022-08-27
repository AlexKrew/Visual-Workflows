package nodes

import (
	"fmt"
	"workflows/shared/shared_entities"
)

type NodeInput struct {
	input shared_entities.JobPayload
}

func NodeInputFromJob(job shared_entities.Job) NodeInput {
	return NodeInput{
		input: job.Input,
	}
}

func (in *NodeInput) ValueFor(key string) (*shared_entities.WorkflowMessage, error) {
	for _, item := range in.input {
		if item.PortIdentifier == key {
			return &item.Value, nil
		}
	}

	return nil, fmt.Errorf("input with identifier `%s` does not exist", key)
}

type NodeOutput struct {
	output      map[string]shared_entities.WorkflowMessage
	groupOutput map[string]map[string]shared_entities.WorkflowMessage
	log         []string
}

func NewNodeOutput() NodeOutput {
	return NodeOutput{
		output:      make(map[string]shared_entities.WorkflowMessage),
		groupOutput: make(map[string]map[string]shared_entities.WorkflowMessage),
		log:         []string{},
	}
}

func (out *NodeOutput) Set(key string, value shared_entities.WorkflowMessage) {
	out.output[key] = value
}

func (out *NodeOutput) SetGroup(group string, key string, value shared_entities.WorkflowMessage) {
	if _, exists := out.groupOutput[group]; !exists {
		out.groupOutput[group] = make(map[string]shared_entities.WorkflowMessage)
	}

	out.groupOutput[group][key] = value
}

func (out *NodeOutput) Log(msg string) {
	out.log = append(out.log, msg)
}

func (out *NodeOutput) GetOutput() map[string]shared_entities.WorkflowMessage {
	return out.output
}

func (out *NodeOutput) GetGroupOutput() map[string]map[string]shared_entities.WorkflowMessage {
	return out.groupOutput
}

func (out *NodeOutput) GetLogs() []string {
	return out.log
}
