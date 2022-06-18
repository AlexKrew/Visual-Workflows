package main

import (
	"fmt"
	"visualWorkflows/internal/container"
	"visualWorkflows/server"
)

func main() {
	fmt.Println("Starting the server...")

	runtimeContainer := container.Construct()

	// availableWorkflows, err := storage.GetAvailableWorkflows()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Available Workflows", availableWorkflows)

	server.StartServer(&runtimeContainer)
}
