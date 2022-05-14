package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/storage"

	"golang.org/x/exp/slices"
)

type Runtime struct {
	Initialized bool

	knownNodes []storage.NodeConfiguration
}

// Initialize prepares the runtime element for handling workflows
func ConstructRuntime() Runtime {

	rt := Runtime{}
	rt.Initialized = false

	fmt.Println("Initializing the runtime ...")

	err := rt.loadKnownNodes()
	if err != nil {
		panic("Failed to initialize runtime")
	}

	rt.Initialized = true
	fmt.Println("Runtime initialized")

	return rt
}

func (rt *Runtime) loadKnownNodes() error {

	knownNodes, err := storage.LoadKnownNodes()
	if err != nil {
		return err
	}

	rt.knownNodes = knownNodes
	return nil
}

// ExecuteWorkflow loads and runs the workflow defined by `workflowID`
func (rt *Runtime) ExecuteWorkflow(workflowID string) {

	if !rt.Initialized {
		panic("Runtime is not initialized")
	}

	// 1. Load the configuration of the workflow
	fmt.Println("Loading workflow definition...")
	config, err := storage.LoadWorkflowDefinition(workflowID)
	if err != nil {
		panic("Failed to load workflow config")
	}

	// 2. Validate the loaded workflow definition
	fmt.Println("Validating workflow definition")
	err = rt.validateWorkflowDefinition(config)
	if err != nil {
		panic("Workflow validation failed")
	}

	// 3. Merge workflow definition with default configuration values
	//
	// TODO:

	// 4. Setup message router
	// router := constructMessageRouter(rt)

	fmt.Println("Ready to run workflow", workflowID)
}

func (rt *Runtime) validateWorkflowDefinition(config storage.WorkflowConfiguration) error {

	if !rt.Initialized {
		return errors.New("runtime not initialized")
	}

	allNodesKnown, err := rt.validateNodeTypes(config.Nodes)
	if err != nil {
		return errors.New("failed to validate node types")
	}
	if !allNodesKnown {
		return errors.New("workflow definition contains unknown node type(s)")
	}

	return nil
}

func (rt *Runtime) validateNodeTypes(nodes map[string]storage.NodeDefinition) (bool, error) {
	nodeTypes := []string{}

	for _, nodeConfig := range rt.knownNodes {
		nodeTypes = append(nodeTypes, nodeConfig.Type)
	}

	for _, nodeDef := range nodes {

		if !slices.Contains(nodeTypes, nodeDef["type"].(string)) {
			return false, nil
		}
	}

	return true, nil
}
