package workflows

import (
	"workflows/shared/shared_entities"
)

type MessageCache struct {
	Workflow *Workflow

	NodeMessageStores map[NodeID]StoredMessages
}

type StoredMessages = map[string]shared_entities.WorkflowMessage

func ConstructMessageCache(workflow *Workflow) (MessageCache, error) {

	stores := make(map[NodeID]StoredMessages)

	for _, node := range workflow.Nodes {

		store := make(map[string]shared_entities.WorkflowMessage)

		ports := node.Ports
		for _, port := range ports {

			if port.IsInputPort {
				portMessage := port.DefaultMessage
				store[port.Identifier] = portMessage
			}
		}

		stores[node.ID] = store
	}

	cache := MessageCache{
		Workflow:          workflow,
		NodeMessageStores: stores,
	}

	return cache, nil
}

func (cache *MessageCache) MessagesForNodeId(nodeId NodeID) (StoredMessages, bool) {
	messages, exists := cache.NodeMessageStores[nodeId]
	return messages, exists
}

func (cache *MessageCache) SetMessage(portAddr PortAddress, message shared_entities.WorkflowMessage) {

	node, exists := cache.Workflow.NodeByID(portAddr.NodeID)
	if !exists {
		panic("node does not exist")
	}

	port, exists := node.PortByID(portAddr.PortID)
	if !exists {
		panic("port does not exist")
	}

	nodePorts := cache.NodeMessageStores[portAddr.NodeID]
	nodePorts[port.Identifier] = message
}
