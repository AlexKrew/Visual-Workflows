package runtime

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

type EventStreamer struct {
	eventSource chan rxgo.Item
	Stream      *rxgo.Observable
	runtime     *Runtime
}

func constructEventStream(runtime *Runtime) (*EventStreamer, error) {
	ch := make(chan rxgo.Item)
	stream := rxgo.FromEventSource(ch)

	return &EventStreamer{
		eventSource: ch,
		Stream:      &stream,
		runtime:     runtime,
	}, nil
}

func (es *EventStreamer) addOperation(op Operation) error {
	op(es.Stream, OperationProps{
		eventStreamer: es,
		messageStore:  es.runtime.Store,
		messageRouter: es.runtime.Router,
		jobQueue:      es.runtime.JobQueue,
		runtimeEvents: es.runtime.Events,
	})

	// may not return when operation is blocking
	return nil
}

func (es *EventStreamer) invokeEvent(event Event) {
	fmt.Println("Invoke EVENT", event.Type)
	es.eventSource <- rxgo.Of(event)
}
