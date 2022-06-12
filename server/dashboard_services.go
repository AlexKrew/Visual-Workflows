package server

import "github.com/gin-gonic/gin"

func registerDashboardServices(rg *gin.RouterGroup) {
	dashboard := rg.Group("/dashboard")
	{
		dashboard.GET("/", getWorkflows)
		dashboard.POST("/", createWorkflow)
		dashboard.POST("/{workflowId}/start", startWorkflow)
		dashboard.POST("/{workflowId}/stop", stopWorkflow)
		dashboard.POST("/{workflowId}/terminate", terminateWorkflow)
		dashboard.DELETE("/{workflowId}", deleteWorkflow)
	}
}

func getWorkflows(c *gin.Context) {

}

func createWorkflow(c *gin.Context) {

}

func startWorkflow(c *gin.Context) {

}

func stopWorkflow(c *gin.Context) {

}

func terminateWorkflow(c *gin.Context) {

}

func deleteWorkflow(c *gin.Context) {

}
