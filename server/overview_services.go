package server

import (
	"net/http"

	container "visualWorkflows/internal/container"
	"visualWorkflows/internal/storage"

	"github.com/gin-gonic/gin"
)

var wfContainer *container.WorkflowContainer

func registerOverviewServices(rg *gin.RouterGroup, container *container.WorkflowContainer) {
	wfContainer = container

	overviewServices := rg.Group("/workflows")

	overviewServices.GET("", getWorkflows)
	overviewServices.POST("new", createWorkflow)
}

func getWorkflows(c *gin.Context) {
	wfInfos, err := wfContainer.GetAvailableWorkflows()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, wfInfos)
}

func createWorkflow(c *gin.Context) {
	var request CreateWorkflowRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	props := storage.CreateWorkflowProps{
		Name: request.Name,
	}

	id, err := wfContainer.CreateWorkflow(props)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
