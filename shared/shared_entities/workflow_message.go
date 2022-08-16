package shared_entities

type MessageType = string

// type MessageType uint8

// const (
// 	EmptyMessageType MessageType = iota
// 	StringMessageType
// 	BooleanMessageType
// 	NumberMessageType
// 	AnyMessageType
// )

var (
	StringMessageType  = "STRING"
	NumberMessageType  = "NUMBER"
	BooleanMessageType = "BOOL"
	AnyMessageType     = "EMPTY"
	EmptyMessageType   = "EMPTY"
)

type WorkflowMessage struct {
	DataType MessageType `json:"datatype"`
	Value    any         `json:"value"`
}

// func MessageTypeFromString(messagetype string) MessageType {
// 	if messagetype == "STRING" {
// 		return StringMessageType
// 	}
// 	if messagetype == "NUMBER" {
// 		return NumberMessageType
// 	}
// 	if messagetype == "BOOL" {
// 		return BooleanMessageType
// 	}
// 	if messagetype == "ANY" {
// 		return AnyMessageType
// 	}

// 	return EmptyMessageType
// }

// func MessageTypeToString(messagetype MessageType) string {
// 	if messagetype == StringMessageType {
// 		return "STRING"
// 	}
// 	if messagetype == NumberMessageType {
// 		return "NUMBER"
// 	}
// 	if messagetype == BooleanMessageType {
// 		return "BOOL"
// 	}
// 	if messagetype == AnyMessageType {
// 		return "ANY"
// 	}

// 	return ""
// }

/* Builder functions */

func EmptyMessage() WorkflowMessage {
	return WorkflowMessage{
		DataType: EmptyMessageType,
	}
}

func StringMessage(body string) WorkflowMessage {
	return WorkflowMessage{
		DataType: StringMessageType,
		Value:    body,
	}
}

func BooleanMessage(body bool) WorkflowMessage {
	return WorkflowMessage{
		DataType: BooleanMessageType,
		Value:    body,
	}
}

func NumberMessage(body int) WorkflowMessage {
	return WorkflowMessage{
		DataType: NumberMessageType,
		Value:    body,
	}
}

func AnyMessage(body any) WorkflowMessage {
	return WorkflowMessage{
		DataType: AnyMessageType,
		Value:    body,
	}
}

/* Helper functions */
func (msg *WorkflowMessage) IsEmpty() bool {
	return msg.DataType == EmptyMessageType
}
