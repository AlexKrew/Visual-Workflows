package webserver

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
		workflows.PATCH("/:id", updateWorkflow)
		workflows.PATCH("/:id/start", startWorkflow)
	}
}

func getWorkflow(c *gin.Context) {
	var workflowId string = c.Param("id")

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
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	err = WFHelper.PublishChanges(workflow)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "Correct")
}

func startWorkflow(c *gin.Context) {
	workflowID := c.Param("id")

	err := WFHelper.StartWorkflow(workflowID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, workflowID)
}
