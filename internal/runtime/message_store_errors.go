package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/entities"
)

func missingNodeCacheKeyError(nodeID entities.NodeID) error {
	msg := fmt.Sprintf("no cache key for node id '%s'", nodeID)
	return errors.New(msg)
}

func missingPortCacheKeyError(nodeID entities.NodeID, portID entities.PortID) error {
	msg := fmt.Sprintf("no cache key for port id '%s' of node %s", portID, nodeID)
	return errors.New(msg)
}
