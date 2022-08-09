package workflows

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"workflows/internal/utils"
)

const (
	pathToWorkflows string = "./workflows"
	pathToNodes     string = "./nodes"
)

var (
	workflowFileSuffix = "vwf.json"
)

func WorkflowFromFilesystem(id utils.UUID) (Workflow, error) {
	filePath := fmt.Sprintf("%s/%s.%s", pathToWorkflows, id, workflowFileSuffix)
	fmt.Println("Loading workflow from file", filePath)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return Workflow{}, errors.New("failed to open file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var workflow Workflow

	json.Unmarshal(byteValue, &workflow)
	return workflow, nil
}

func WorkflowToFilesystem(workflow Workflow) error {
	filePath := fmt.Sprintf("%s/%s.test.vwf.json", pathToWorkflows, workflow.ID)

	file, _ := json.MarshalIndent(workflow, "", " ")

	err := ioutil.WriteFile(filePath, file, 0644)

	return err
}

type WorkflowInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AvailableWorkflows() ([]WorkflowInfo, error) {

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
