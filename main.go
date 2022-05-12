package main

import (
	"fmt"
	"visualWorkflows/internal/runtime"
)

func main() {
	fmt.Println("Hello from main")

	runtime := runtime.Runtime{}
	runtime.Initialize()

	runtime.ExecuteWorkflow("flow1")
}
