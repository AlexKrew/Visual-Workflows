package server

import (
	"fmt"
	"net/http"
	container "visualWorkflows/internal/container"

	"github.com/gin-gonic/gin"
)

func registerWorkflowServices(rg *gin.RouterGroup, container *container.WorkflowContainer) {
	wfContainer = container

	workflows := rg.Group("/workflows")
	{
		workflows.GET("/:id", getWorkflow)
		workflows.PATCH("/:id", updateWorkflow)
	}
}

func getWorkflow(c *gin.Context) {
	var workflowId string = c.Param("id")
	fmt.Println("GET WORKFLOW", workflowId)

	workflow, err := wfContainer.GetWorkflowById(workflowId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// return c.JSON(http.StatusOK, workflow)
	c.JSON(http.StatusOK, gin.H{"workflow": workflow})
	return
}

func updateWorkflow(c *gin.Context) {

}
