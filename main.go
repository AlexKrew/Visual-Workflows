package main

import (
	"fmt"
	"visualWorkflows/server"
	// "visualWorkflows/internal/runtime"
)

func main() {
	fmt.Println("Hello from main")

	server.StartServer()

	// runtime := runtime.ConstructRuntime()
	// runtime.ExecuteWorkflow("flow1")
}
