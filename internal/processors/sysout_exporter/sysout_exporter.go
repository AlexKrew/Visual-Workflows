package sysout_exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"workflows/internal/workflows"

	"github.com/reactivex/rxgo/v2"
)

type SysoutExporter struct {
	outfile *os.File
}

func ConstructSysoutExporter(outfilePath string) (*SysoutExporter, error) {
	f, err := os.OpenFile(outfilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	exporter := SysoutExporter{
		outfile: f,
	}

	return &exporter, nil
}

func (exporter *SysoutExporter) Register(eventStream *workflows.EventStream) {
	// Register commands handler
	go exporter.registerCommandsHandler(eventStream.CommandsObservable)
	go exporter.registerEventsHandler(eventStream.EventsObservable)

	fmt.Println("Registered SysoutExporter")
}

func (exporter *SysoutExporter) registerCommandsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		exporter.handleCommand(i.(workflows.WorkflowCommand))
	}, func(err error) {
		fmt.Println("Error", err)
	}, func() {})
}

func (exporter *SysoutExporter) registerEventsHandler(observable *rxgo.Observable) {
	(*observable).ForEach(func(i interface{}) {
		exporter.handleEvent(i.(workflows.WorkflowEvent))
	}, func(err error) {
		fmt.Println("Error", err)
	}, func() {})
}

func (exporter *SysoutExporter) handleCommand(command workflows.WorkflowCommand) {

	commandJSON, err := json.Marshal(command)
	commandJSON = append(commandJSON, '\n')
	if err != nil {
		panic(err)
	}

	exporter.outfile.Write(commandJSON)
}

func (exporter *SysoutExporter) handleEvent(event workflows.WorkflowEvent) {

	eventJSON, err := json.Marshal(event)
	eventJSON = append(eventJSON, '\n')
	if err != nil {
		panic(err)
	}

	_, err = exporter.outfile.Write(eventJSON)
	if err != nil {
		panic(err)
	}
}
