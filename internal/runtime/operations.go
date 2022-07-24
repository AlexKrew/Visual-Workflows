package runtime

import (
	"github.com/reactivex/rxgo/v2"
)

type OperationProps struct {
	eventStreamer *EventStreamer
	messageStore  *MessageStore
	messageRouter *MessageRouter
	jobQueue      *JobQueue
	runtimeEvents chan interface{}
}
type Operation func(*rxgo.Observable, OperationProps)
