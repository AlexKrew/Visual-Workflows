package main

import (
	"fmt"
	"visualWorkflows/internal/runtime"
)

func main() {
	fmt.Println("Hello from main")

	runtime.StartWorkflow("flow1")
}
