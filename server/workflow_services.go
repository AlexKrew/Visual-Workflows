package server

import "github.com/gin-gonic/gin"

func registerWorkflowServices(rg *gin.RouterGroup) {
	workflows := rg.Group("/workflows/{workflow-id}")
	{
		workflows.GET("/", getWorkflow)
		workflows.PATCH("/", updateWorkflow)
	}
}

func getWorkflow(c *gin.Context) {

}

func updateWorkflow(c *gin.Context) {

}
