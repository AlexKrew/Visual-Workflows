package workflows

import (
	"log"

	"github.com/reactivex/rxgo/v2"
)

type EventStream struct {
	CommandChannel chan rxgo.Item
	EventChannel   chan rxgo.Item

	CommandsObservable *rxgo.Observable
	EventsObservable   *rxgo.Observable
}

func ConstructEventStream() *EventStream {
	commandChannel := make(chan rxgo.Item, 10)
	commandsObservable := rxgo.FromEventSource(commandChannel)

	eventChannel := make(chan rxgo.Item, 10)
	eventsObservable := rxgo.FromEventSource(eventChannel)

	return &EventStream{
		CommandChannel:     commandChannel,
		EventChannel:       eventChannel,
		CommandsObservable: &commandsObservable,
		EventsObservable:   &eventsObservable,
	}
}

func (eventStream *EventStream) AddCommand(command WorkflowCommand) {
	log.Println("Add command", command.Type)
	eventStream.CommandChannel <- rxgo.Of(command)
	log.Println("Added command", command.Type)
}

func (eventStream *EventStream) AddEvent(event WorkflowEvent) {
	log.Println("Add event", event.Type)
	eventStream.EventChannel <- rxgo.Of(event)
	log.Println("Added event", event.Type)
}
