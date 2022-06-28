package storage

import "fmt"

type CreateWorkflowProps struct {
	Name string
}

func CreateWorkflow(props CreateWorkflowProps) (string, error) {
	fmt.Println("Create workflow", props)
	return "1", nil
}
