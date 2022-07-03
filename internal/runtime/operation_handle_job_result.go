package runtime

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func handleJobResultOperation(observable *rxgo.Observable, props OperationProps) {

	<-(*observable).
		Filter(func(item interface{}) bool {
			return item.(Event).Type == ETJobResultReceived
		}).
		ForEach(func(event interface{}) {

			fmt.Println("----- Received", event)
			body := event.(Event).Body.(JobResultReceivedBody)
			props.runtimeEvents <- body.Result

		}, func(err error) {}, func() {})
}
