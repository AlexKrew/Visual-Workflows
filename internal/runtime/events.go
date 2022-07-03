package runtime

import (
	"visualWorkflows/internal/entities"

	"github.com/google/uuid"
)

type EventID = uuid.UUID
type EventType = int16

const (
	ETMessagesReceived EventType = iota
)

type Event struct {
	ID   EventID
	Type EventType
	Body any
}

func buildNewEvent(body any, t EventType) Event {
	return Event{
		ID:   getNewUUID(),
		Type: t,
		Body: body,
	}
}

/* MessagesReceived Event */

type MessagesReceivedBody struct {
	nodeId entities.NodeID
}

func buildMessagesReceivedEvent(body MessagesReceivedBody) Event {
	return buildNewEvent(body, ETMessagesReceived)
}

/* ExampleEvent */

// type ExampleEventBody struct {
// 	id entities.NodeID
// }

// func buildExampleEvent(body ExampleEventBody) Event {
// 	return buildNewEvent(body, type)
// }
