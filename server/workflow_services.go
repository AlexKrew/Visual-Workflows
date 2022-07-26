package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	container "visualWorkflows/internal/container"
	"visualWorkflows/shared/entities"

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
}

func updateWorkflow(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	var workflow entities.Workflow
	json.Unmarshal(jsonData, &workflow)
	fmt.Println("WORKFLOW", workflow)
	c.String(http.StatusOK, "Correct")
}
