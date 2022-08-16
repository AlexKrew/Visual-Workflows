package workflows

type MessageType uint8

const (
	EmptyMessageType MessageType = iota
	StringMessageType
	BooleanMessageType
	NumberMessageType
	AnyMessageType
)

type Message struct {
	DataType MessageType `json:"datatype"`
	Value    any         `json:"value"`
}

func MessageTypeFromString(messagetype string) MessageType {
	if messagetype == "STRING" {
		return StringMessageType
	}
	if messagetype == "NUMBER" {
		return NumberMessageType
	}
	if messagetype == "BOOL" {
		return BooleanMessageType
	}
	if messagetype == "ANY" {
		return AnyMessageType
	}

	return EmptyMessageType
}

func MessageTypeToString(messagetype MessageType) string {
	if messagetype == StringMessageType {
		return "STRING"
	}
	if messagetype == NumberMessageType {
		return "NUMBER"
	}
	if messagetype == BooleanMessageType {
		return "BOOL"
	}
	if messagetype == AnyMessageType {
		return "ANY"
	}

	return ""
}

/* Builder functions */

func EmptyMessage() Message {
	return Message{
		DataType: EmptyMessageType,
	}
}

func StringMessage(body string) Message {
	return Message{
		DataType: StringMessageType,
		Value:    body,
	}
}

func BooleanMessage(body bool) Message {
	return Message{
		DataType: BooleanMessageType,
		Value:    body,
	}
}

func NumberMessage(body int) Message {
	return Message{
		DataType: NumberMessageType,
		Value:    body,
	}
}

func AnyMessage(body any) Message {
	return Message{
		DataType: AnyMessageType,
		Value:    body,
	}
}

/* Helper functions */
func (msg *Message) IsEmpty() bool {
	return msg.DataType == EmptyMessageType
}