package entities

type Input struct {
	Messages map[string]any
}

func (input *Input) GetMessage(portID PortID) WorkflowMessage {
	// check if msg at port is a variable
	return WorkflowMessage{}
}

func (input *Input) GetGroupMessages(groupPortID GroupPortID) map[string]WorkflowMessage {
	return nil
}

type Output struct {
}
