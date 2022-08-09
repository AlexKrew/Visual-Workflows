package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"workflows/internal/workflows"

	"github.com/gin-gonic/gin"
)

func registerWorkflowServices(rg *gin.RouterGroup) {
	workflows := rg.Group("/workflows")
	{
		workflows.GET("/:id", getWorkflow)
		workflows.GET("/:id/start", startWorkflow)
		workflows.PATCH("/:id", updateWorkflow)
	}
}

func getWorkflow(c *gin.Context) {
	var workflowId string = c.Param("id")
	fmt.Println("GET WORKFLOW", workflowId)

	workflow, exists := WFHelper.WorkflowById(workflowId)
	if !exists {
		c.String(http.StatusBadRequest, errors.New("workflow does not exist").Error())
		return
	}
	// return c.JSON(http.StatusOK, workflow)
	c.JSON(http.StatusOK, gin.H{"workflow": workflow})
}

func updateWorkflow(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	var workflow workflows.Workflow
	err = json.Unmarshal(jsonData, &workflow)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	err = WFHelper.PublishChanges(workflow)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	c.String(http.StatusOK, "Correct")
}

func startWorkflow(c *gin.Context) {
	workflowID := c.Param("id")

	err := WFHelper.workflowProcessor.StartWorkflow(workflowID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, workflowID)
}
