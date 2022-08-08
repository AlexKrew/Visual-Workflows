package workflows

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"workflows/internal/utils"
)

const (
	pathToWorkflows string = "/Users/mfa/code/master/project-2/engine/example-workflows"
	pathToNodes     string = "/workspaces/Visual-Workflows/nodes"
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
