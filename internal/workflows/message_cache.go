package workflows

import (
	"log"
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

func (cache *MessageCache) JobPayloadForNodeId(nodeId NodeID) (shared_entities.JobPayload, bool) {

	node, nodeExists := cache.Workflow.NodeByID(nodeId)
	if !nodeExists {
		return nil, false
	}

	payload := shared_entities.JobPayload{}

	for _, port := range node.Ports {

		if !port.IsInputPort {
			continue
		}

		message, exists := cache.GetMessage(PortAddress{
			NodeID: node.ID,
			PortID: port.ID,
		})
		if !exists {
			log.Panicf("missing port value in cache for id `%s`", port.ID)
			return nil, false
		}

		payload = append(payload, shared_entities.JobPayloadItem{
			NodeID:         node.ID,
			PortIdentifier: port.Identifier,
			GroupID:        port.GroupID,
			Value:          message,
		})
	}

	return payload, true
}

func (cache *MessageCache) GetMessage(portAddr PortAddress) (shared_entities.WorkflowMessage, bool) {
	nodeMessages, exists := cache.NodeMessageStores[portAddr.NodeID]
	if !exists {
		return shared_entities.EmptyMessage(), false
	}

	for portId, message := range nodeMessages {
		if portAddr.PortID == portId {
			return message, true
		}
	}

	return shared_entities.EmptyMessage(), false
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
	nodePorts[port.ID] = message
}
