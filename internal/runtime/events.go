package runtime

import (
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/utils"

	"github.com/google/uuid"
)

type EventID = uuid.UUID
type EventType = int16

const (
	ETMessagesReceived EventType = iota
	ETJobResultReceived
)

type Event struct {
	ID   EventID
	Type EventType
	Body any
}

func buildNewEvent(body any, t EventType) Event {
	return Event{
		ID:   utils.GetNewUUID(),
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

/* JobResultReceived Event */
type JobResultReceivedBody struct {
	Result entities.JobResult
}

func buildJobResultReceivedEvent(body JobResultReceivedBody) Event {
	return buildNewEvent(body, ETJobResultReceived)
}

/* ExampleEvent */

// type ExampleEventBody struct {
// 	id entities.NodeID
// }

// func buildExampleEvent(body ExampleEventBody) Event {
// 	return buildNewEvent(body, type)
// }
