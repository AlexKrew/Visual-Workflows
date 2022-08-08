package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerOverviewServices(rg *gin.RouterGroup) {

	overviewServices := rg.Group("/workflows")

	overviewServices.GET("", getWorkflows)
	overviewServices.POST("new", createWorkflow)
}

func getWorkflows(c *gin.Context) {
	// wfInfos, err := wfContainer.GetAvailableWorkflows()
	// if err != nil {
	// 	panic(err)
	// }

	// c.JSON(http.StatusOK, wfInfos)
	c.JSON(http.StatusOK, nil)
}

func createWorkflow(c *gin.Context) {
	var request CreateWorkflowRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// WorkflowHelper.CreateNewWorkflow(request.Name)

	c.JSON(http.StatusCreated, gin.H{"status": "okay"})
}
