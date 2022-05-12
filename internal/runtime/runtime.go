package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/storage"

	"golang.org/x/exp/slices"
)

type Runtime struct {
	Initialized bool

	NodeTypes []storage.NodeConfig
}

// Initialize prepares the runtime element for handling workflows
func (rt *Runtime) Initialize() {

	fmt.Println("Initializing the runtime ...")

	// 1. Load available node types that can be executed by the runtime
	availableNodeTypes, err := storage.LoadAvailableNodes()
	if err != nil {
		panic("Failed to load available nodes")
	}
	rt.NodeTypes = availableNodeTypes

	rt.Initialized = true
	fmt.Println("Runtime initialized")
}

func (rt *Runtime) ExecuteWorkflow(workflowID string) {

	if !rt.Initialized {
		panic("Runtime is not initialized")
	}

	// 1. Load the configuration of the workflow
	fmt.Println("Loading workflow...")
	config, err := storage.LoadWorkflowConfig(workflowID)
	if err != nil {
		panic("Failed to load workflow config")
	}

	// 2. Validate the loaded config
	err = rt.validateWorkflowConfig(config)
	if err != nil {
		panic("Workflow validation failed")
	}
	fmt.Println("Workflow validated")
}

func (rt *Runtime) validateWorkflowConfig(config storage.WorkflowConfiguration) error {

	// check if the type of nodes are known to the runtime
	availableNodeTypes, err := rt.availableNodeTypes()
	if err != nil {
		panic(err)
	}

	for _, nodeConfig := range config.Nodes {
		if t, ok := nodeConfig["type"].(string); ok {

			if !slices.Contains(availableNodeTypes, t) {
				fmt.Println(t)
				panic("Unknown node type")
			}

		} else {
			return errors.New("missing type")
		}

	}

	return nil
}

func (rt *Runtime) availableNodeTypes() ([]string, error) {

	if !rt.Initialized {
		return []string{}, errors.New("Runtime is not initialized")
	}

	types := []string{}

	fmt.Println("Loading types:")
	for _, nodeConfig := range rt.NodeTypes {
		fmt.Println(nodeConfig.Type)
		types = append(types, nodeConfig.Type)
	}

	return types, nil
}
