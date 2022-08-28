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

// ValueFor returns a message for the given identifier if a value exists.
// Only use this if there is surly only one port with this identifier (i.e. not a group port)
// In this case use the ValueForGroup function
func (in *NodeInput) ValueFor(key string) (*shared_entities.WorkflowMessage, error) {
	for _, item := range in.input {
		if item.PortIdentifier == key {
			return &item.Value, nil
		}
	}

	return nil, fmt.Errorf("input with identifier `%s` does not exist", key)
}

func (in *NodeInput) Groups() []string {
	groups := []string{}

	for _, item := range in.input {
		if item.GroupID != "" {
			groups = append(groups, item.GroupID)
		}
	}

	return groups
}

func (in *NodeInput) ValueForGroup(group string, key string) (*shared_entities.WorkflowMessage, error) {
	for _, item := range in.input {
		if item.GroupID == group && item.PortIdentifier == key {
			return &item.Value, nil
		}
	}

	return nil, fmt.Errorf("input with identifier `%s` for group `%s` does not exist", key, group)
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
