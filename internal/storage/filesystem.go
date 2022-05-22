package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"visualWorkflows/internal/entities"
)

var pathToWorkflows string = "/workspaces/Visual-Workflows/workflows"
var pathToNodes string = "/workspaces/Visual-Workflows/nodes"

func LoadWorkflowDefinition(workflowID string) (entities.Workflow, error) {
	fmt.Println("Reading the workflow from file")

	filePath := fmt.Sprintf("%s/%s.vwf.json", pathToWorkflows, workflowID)
	fmt.Println("Open", filePath)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file")
		return entities.Workflow{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config entities.Workflow

	json.Unmarshal(byteValue, &config)
	fmt.Println("Loaded config of flow", config.Name)

	return config, nil
}

func LoadKnownNodes() ([]entities.Node, error) {
	// TODO: Load all .vwf-node.json files from folder
	files := []string{"debug", "inject"}

	nodes := []entities.Node{}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s.vwf-node.json", pathToNodes, file)

		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Failed to open node file")
			return []entities.Node{}, err
		}
		defer jsonFile.Close()

		node := loadNode(jsonFile)
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func loadNode(jsonFile *os.File) entities.Node {
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var node entities.Node

	json.Unmarshal(byteValue, &node)
	return node
}
