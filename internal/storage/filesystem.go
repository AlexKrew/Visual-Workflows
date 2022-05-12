package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var pathToWorkflows string = "/workspaces/Visual-Workflows/workflows"
var pathToNodes string = "/workspaces/Visual-Workflows/nodes"

func LoadWorkflowConfig(workflowID string) (WorkflowConfiguration, error) {
	fmt.Println("Reading the workflow from file")

	filePath := fmt.Sprintf("%s/%s.vwf.json", pathToWorkflows, workflowID)
	fmt.Println("Open", filePath)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file")
		return WorkflowConfiguration{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config WorkflowConfiguration

	json.Unmarshal(byteValue, &config)
	fmt.Println("Loaded config of flow", config.Name)

	return config, nil
}

func LoadAvailableNodes() ([]NodeConfig, error) {
	// TODO: Load all .vwf-node.json files from folder
	files := []string{"debug", "inject"}

	nodeConfigs := []NodeConfig{}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s.vwf-node.json", pathToNodes, file)

		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Failed to open node file")
			return []NodeConfig{}, err
		}
		defer jsonFile.Close()

		nodeConfig := loadNodeConfig(jsonFile)
		nodeConfigs = append(nodeConfigs, nodeConfig)
	}

	return nodeConfigs, nil
}

func loadNodeConfig(jsonFile *os.File) NodeConfig {
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config NodeConfig

	json.Unmarshal(byteValue, &config)
	return config
}
