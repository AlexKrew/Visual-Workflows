package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LoadWorkflowProps struct {
	ID string
}

func LoadWorkflowDefinition(props LoadWorkflowProps) (WorkflowDefinition, error) {

	filePath := fmt.Sprintf("%s/%s.vwf.json", pathToWorkflows, props.ID)
	fmt.Println("Reading the workflow from file at", filePath)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file")
		return WorkflowDefinition{}, err
	}
	defer jsonFile.Close()

	fmt.Println("Reading the workflow definition")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var definition WorkflowDefinition

	json.Unmarshal(byteValue, &definition)
	fmt.Println("Loaded definition", definition.Name)

	return definition, nil
}
