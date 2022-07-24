package runtime

import (
	"visualWorkflows/shared/entities"

	"github.com/reactivex/rxgo/v2"
)

func handleJobResultOperation(observable *rxgo.Observable, props OperationProps) {

	<-(*observable).
		Filter(func(item interface{}) bool {
			return item.(Event).Type == ETJobResultReceived
		}).
		ForEach(func(event interface{}) {

			body := event.(Event).Body.(JobResultReceivedBody)
			// props.runtimeEvents <- body.Result

			// Give output to message store
			for portId, msg := range body.Result.Output {
				addr := entities.PortAddress{
					NodeID: body.Result.NodeId,
					PortID: portId,
				}
				err := props.messageRouter.publishMessage(addr.GetUniquePortID(), msg)

				if err != nil {
					panic(err)
				}
			}

		}, func(err error) {}, func() {})
}
