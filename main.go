package main

import (
	"fmt"
	"visualWorkflows/internal/runtime"
)

func main() {
	fmt.Println("Hello from main")

	runtime := runtime.ConstructRuntime()

	runtime.ExecuteWorkflow("flow1")
}
