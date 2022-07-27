package main

import (
	"fmt"
	"sync"
	"visualWorkflows/internal/container"
	"visualWorkflows/internal/storage"
	"visualWorkflows/server"
	wc "visualWorkflows/workerclient"
)

var wg sync.WaitGroup

func main() {
	wg.Add(5)

	fmt.Println("Starting the server...")

	runtimeContainer := container.Construct()

	go server.StartServer(&runtimeContainer)

	err := runtimeContainer.LoadWorkflow(storage.LoadWorkflowProps{
		ID: "flow1",
	})
	err = runtimeContainer.LoadWorkflow(storage.LoadWorkflowProps{
		ID: "testflow",
	})

	if err != nil {
		panic(err)
	}

	worker1 := wc.ConstructWorker()
	runtimeContainer.RegisterWorker("flow1", worker1)

	go runtimeContainer.StartWorkflow("flow1")

	wg.Wait()

	// availableWorkflows, err := storage.GetAvailableWorkflows()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Available Workflows", availableWorkflows)

	// server.StartServer(&runtimeContainer)
}
