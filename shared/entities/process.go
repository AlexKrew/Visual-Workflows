package entities

import "fmt"

type Input struct {
	Messages map[PortID]WorkflowMessage
}

func (input *Input) GetMessage(portID PortID) WorkflowMessage {
	msg, exists := input.Messages[portID]
	if !exists {
		fmt.Println("missing msg for port", portID)
		return EmptyMessage()
	}

	return msg
}

func (input *Input) GetGroupMessages(groupPortID GroupPortID) map[PortID]WorkflowMessage {
	return nil
}

type Output struct {
	Messages map[PortID]WorkflowMessage
}

func (output *Output) Add(portID PortID, msg WorkflowMessage) {
	output.Messages[portID] = msg
}
