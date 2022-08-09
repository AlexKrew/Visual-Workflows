package workflows

import "fmt"

type MessageCache struct {
	Workflow *Workflow

	NodeMessageStores map[NodeID]StoredMessages
}

type StoredMessages = map[PortID]Message

func ConstructMessageCache(workflow *Workflow) (MessageCache, error) {

	stores := make(map[NodeID]StoredMessages)

	for _, node := range workflow.Nodes {

		store := make(map[string]Message)

		ports := node.Ports
		for _, port := range ports {

			if port.IsInputPort {
				portMessage := port.DefaultMessage
				store[port.ID] = portMessage
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

func (cache *MessageCache) SetMessage(portAddr PortAddress, message Message) {

	nodePorts := cache.NodeMessageStores[portAddr.NodeID]
	nodePorts[portAddr.PortID] = message

	fmt.Println("Updated message for ", portAddr)
}
