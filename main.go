package main

import (
	"fmt"
	"visualWorkflows/server"
	// "visualWorkflows/internal/runtime"
)

func main() {
	fmt.Println("Starting the server...")
	server.StartServer()

	// runtime := runtime.ConstructRuntime()
	// runtime.ExecuteWorkflow("flow1")
}
