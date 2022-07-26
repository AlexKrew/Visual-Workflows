package entities

type MessageType uint8

const (
	MTEmpty MessageType = iota
	MTString
	MTBoolean
	MTNumber
)

type WorkflowMessage struct {
	DataType MessageType `json:"datatype"`
	Value    any         `json:"value"`
}

/* Builder functions */

func EmptyMessage() WorkflowMessage {
	return WorkflowMessage{
		DataType: MTEmpty,
	}
}

func StringMessage(body string) WorkflowMessage {
	return WorkflowMessage{
		DataType: MTString,
		Value:    body,
	}
}

func BooleanMessage(body bool) WorkflowMessage {
	return WorkflowMessage{
		DataType: MTBoolean,
		Value:    body,
	}
}

func NumberMessage(body int) WorkflowMessage {
	return WorkflowMessage{
		DataType: MTNumber,
		Value:    body,
	}
}

/* Helper functions */
func (msg *WorkflowMessage) IsEmpty() bool {
	return msg.DataType == MTEmpty
}
