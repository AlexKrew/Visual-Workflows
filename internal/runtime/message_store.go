package runtime

import (
	"visualWorkflows/internal/entities"
)

// MessageStore holds the latest messages for each InputPort of each Node in a Workflow.
type MessageStore struct {
	cachedMessages map[entities.NodeID](map[entities.PortID]entities.WorkflowMessage)
}

// constructMessageStore uses the Nodes and Ports in a Workflow to initialize
// a new MessageStore with initial messages defined in Workflow.
func constructMessageStore(workflow entities.Workflow) (MessageStore, error) {

	initialMessages := make(map[entities.NodeID](map[entities.PortID]entities.WorkflowMessage))

	for nodeId, node := range workflow.Nodes {

		inputPortIds := node.GetInputPortIds()
		initialMessages[nodeId] = make(map[entities.PortID]entities.WorkflowMessage)

		for _, portId := range inputPortIds {

			port, _ := node.GetInputPort(portId)
			message := port.GetMessage()
			initialMessages[nodeId][portId] = message

		}
	}

	mc := MessageStore{
		cachedMessages: initialMessages,
	}

	return mc, nil
}

func (mc *MessageStore) setMessage(address entities.PortAddress, message entities.WorkflowMessage) error {

	nodeMessages, exists := mc.cachedMessages[address.NodeID]
	if !exists {
		return missingNodeCacheKeyError(address.NodeID)
	}

	_, exists = nodeMessages[address.PortID]
	if !exists {
		return missingPortCacheKeyError(address.NodeID, address.PortID)
	}

	nodeMessages[address.PortID] = message

	return nil
}

func (mc *MessageStore) StoreNewMessage(address entities.PortAddress, message entities.WorkflowMessage) error {

	err := mc.setMessage(address, message)
	if err != nil {
		return err
	}

	// Publish message

	return nil
}
