package runtime

import (
	"fmt"
	"visualWorkflows/shared/entities"

	"github.com/reactivex/rxgo/v2"
)

func createJobOperation(observable *rxgo.Observable, props OperationProps) {

	<-(*observable).
		Filter(func(item interface{}) bool {
			return item.(Event).Type == ETMessagesReceived
		}).
		ForEach(func(event interface{}) {

			body := event.(Event).Body.(MessagesReceivedBody)
			messages, err := props.messageStore.GetMessagesFor(body.nodeId)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Messages for", body.nodeId, messages)
			if !hasEmptyMessage(messages) {

				// TODO: Find a better way
				node := props.eventStreamer.runtime.Workflow.Nodes[body.nodeId]

				jobProps := entities.ProcessJobProps{
					NodeID:   body.nodeId,
					NodeType: node.Type,
					Input:    messages,
				}
				props.jobQueue.addJob(entities.BuildProcessJob(jobProps))
			}

		}, func(err error) {}, func() {})
}

func hasEmptyMessage(messages map[string]entities.WorkflowMessage) bool {
	for _, message := range messages {
		if message.IsEmpty() {
			return true
		}
	}
	return false
}
