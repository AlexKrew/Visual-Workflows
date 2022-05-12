package runtime

import (
	"fmt"
	"visualWorkflows/internal/storage"
)

func StartWorkflow(workflowID string) {
	fmt.Println("Starting Workflow...")

	availableNodeTypes, err := storage.LoadAvailableNodes()
	if err != nil {
		panic("Failed to load available nodes")
	}

	fmt.Println("Available nodes:")
	for _, nodeConfig := range availableNodeTypes {
		fmt.Println(nodeConfig.Name)
	}

	fmt.Println("Loading workflow...")
	config, err := storage.LoadWorkflowConfig(workflowID)
	if err != nil {
		panic("Failed to load workflow config")
	}

	fmt.Println(config.ID, config.Name, config.Edges)
}
