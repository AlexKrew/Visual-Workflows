package server

import "github.com/gin-gonic/gin"

func registerDashboardServices(rg *gin.RouterGroup) {
	dashboard := rg.Group("/dashboard")
	{
		dashboard.GET("/", getWorkflows)
	}
}

func getWorkflows(c *gin.Context) {

}
