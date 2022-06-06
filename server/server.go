package server

import (
	"net/http"

	docs "visualWorkflows/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Visual Workflow API
// @version 0.1
// @description description

// @host localhost:8080
// @BasePath /api/v1
func StartServer() {

	router := gin.Default()

	v1 := router.Group("/api/v1")

	// Register services / routes
	registerHealthServices(v1)
	registerDashboardServices(v1)
	registerWorkflowServices(v1)

	// Setup swagger api documentation
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.ListenAndServe(":8080", router)
}
