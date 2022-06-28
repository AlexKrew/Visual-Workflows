package main

import (
	"fmt"
	"visualWorkflows/internal/container"
	"visualWorkflows/internal/storage"
)

func main() {
	fmt.Println("Starting the server...")

	runtimeContainer := container.Construct()
	err := runtimeContainer.LoadWorkflow(storage.LoadWorkflowProps{
		ID: "flow1",
	})

	if err != nil {
		panic(err)
	}

	// availableWorkflows, err := storage.GetAvailableWorkflows()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Available Workflows", availableWorkflows)

	// server.StartServer(&runtimeContainer)
}
