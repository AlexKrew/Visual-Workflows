package storage

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const (
	// 	pathToWorkflowTemplates string = "/workspaces/Visual-Workflows/workflow-templates"
	pathToWorkflows string = "/workspaces/Visual-Workflows/workflows"
	pathToNodes     string = "/workspaces/Visual-Workflows/nodes"
)

var (
	workflowFileSuffix = ".vwf.json"
)

// GetAvailableWorkflows reads all '.vwf.json' files inside the [pathToWorkflowTemplates] directory.
func GetAvailableWorkflows() ([]WorkflowInfo, error) {

	allFiles, err := os.ReadDir(pathToWorkflows)
	if err != nil {
		return []WorkflowInfo{}, err
	}

	infos := []WorkflowInfo{}

	for _, entry := range allFiles {

		if !isWorkflowFile(entry) {
			continue
		}

		filename := fmt.Sprintf("%s/%s", pathToWorkflows, entry.Name())
		workflowInfo, err := extractWorkflowInformation(filename)
		if err != nil {
			return []WorkflowInfo{}, err
		}

		infos = append(infos, workflowInfo)
	}

	return infos, nil
}

func isWorkflowFile(entry fs.DirEntry) bool {

	if entry.IsDir() {
		return false
	}

	if !strings.HasSuffix(entry.Name(), workflowFileSuffix) {
		return false
	}

	return true
}

func extractWorkflowInformation(filename string) (WorkflowInfo, error) {

	var info WorkflowInfo

	byteValue, err := os.ReadFile(filename)
	if err != nil {
		return info, err
	}

	json.Unmarshal(byteValue, &info)

	return info, nil
}

// func LoadKnownNodes() ([]entities.Node, error) {

// 	// TODO: Load all .vwf-node.json files from folder
// 	files := []string{"debug", "inject"}

// 	nodes := []entities.Node{}

// 	for _, file := range files {
// 		filePath := fmt.Sprintf("%s/%s.vwf-node.json", pathToNodes, file)

// 		jsonFile, err := os.Open(filePath)
// 		if err != nil {
// 			fmt.Println("Failed to open node file")
// 			return []entities.Node{}, err
// 		}
// 		defer jsonFile.Close()

// 		node := loadNode(jsonFile)
// 		nodes = append(nodes, node)
// 	}

// 	return nodes, nil
// }

// func loadNode(jsonFile *os.File) entities.Node {
// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var node entities.Node

// 	json.Unmarshal(byteValue, &node)
// 	return node
// }
