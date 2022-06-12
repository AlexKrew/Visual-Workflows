package runtime

import (
	"errors"
	"fmt"
	"visualWorkflows/internal/entities"
	"visualWorkflows/internal/storage"

	"golang.org/x/exp/slices"
)

type Runtime struct {
	Initialized bool
	knownNodes  []entities.Node

	Workflow entities.Workflow
}

// Initialize prepares the runtime element for handling workflows
func Construct() (Runtime, error) {

	rt := Runtime{}
	rt.Initialized = false

	fmt.Println("Initializing the runtime ...")

	err := rt.loadKnownNodes()
	if err != nil {
		return Runtime{}, err
	}

	rt.Initialized = true
	fmt.Println("Runtime initialized")

	return rt, nil
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
	workflow, err := storage.LoadWorkflowDefinition(workflowID)
	if err != nil {
		panic("Failed to load workflow config")
	}

	// 2. Validate the loaded workflow definition
	fmt.Println("Validating workflow definition")
	err = rt.validateWorkflowDefinition(workflow)
	if err != nil {
		panic(err)
	}

	// Workflow is valid
	rt.Workflow = workflow

	// 3. Construct the message router
	router, err := constructMessageRouter(rt)
	if err != nil {
		panic(err)
	}

	fmt.Println("Ready to run workflow", workflowID, router)
}

func (rt *Runtime) validateWorkflowDefinition(config entities.Workflow) error {

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

func (rt *Runtime) validateNodeTypes(nodes map[string]entities.Node) (bool, error) {
	nodeTypes := []string{}

	for _, nodeConfig := range rt.knownNodes {
		nodeTypes = append(nodeTypes, nodeConfig.Type)
	}

	for _, nodeDef := range nodes {

		if !slices.Contains(nodeTypes, nodeDef.Type) {
			return false, nil
		}
	}

	return true, nil
}
