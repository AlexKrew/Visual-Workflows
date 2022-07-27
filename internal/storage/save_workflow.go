package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveWorkflow(workflowId string, workflowDefinition WorkflowDefinition) error {

	filePath := fmt.Sprintf("%s/%s.test.vwf.json", pathToWorkflows, workflowId)

	file, _ := json.MarshalIndent(workflowDefinition, "", " ")

	err := ioutil.WriteFile(filePath, file, 0644)

	return err
}
