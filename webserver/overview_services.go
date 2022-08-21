package webserver

import (
	"log"
	"net/http"
	"workflows/internal/workflows"

	"github.com/gin-gonic/gin"
)

func registerOverviewServices(rg *gin.RouterGroup) {

	overviewServices := rg.Group("/workflows")

	overviewServices.GET("", getWorkflows)
	overviewServices.POST("", createWorkflow)
}

func getWorkflows(c *gin.Context) {
	wfInfos, err := workflows.AvailableWorkflows()
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

	workflowName := request.Name
	workflow := workflows.NewWorkflow(workflowName)

	err := workflows.WorkflowToFilesystem(workflow)
	if err != nil {
		log.Panicf("Failed to write new workflow to file: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "okay"})
}
