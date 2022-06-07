package server

import (
	"net/http"

	docs "visualWorkflows/docs"

	"github.com/gin-contrib/cors"
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

	// Cors setup
	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
		},
	))

	v1 := router.Group("/api/v1")

	v1.GET("/workflows", func(c *gin.Context) {
		c.String(http.StatusAccepted, "Success")
	})

	// Register services / routes
	registerHealthServices(v1)
	registerDashboardServices(v1)

	// Setup swagger api documentation
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	http.ListenAndServe(":8000", router)
}
