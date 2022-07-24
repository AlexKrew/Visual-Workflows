package server

import (
	"fmt"
	"net/http"

	docs "visualWorkflows/docs"
	wf "visualWorkflows/internal/container"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Visual Workflow API
// @version 0.1
// @description description

// @host localhost:8000
// @BasePath /api/v1
func StartServer(runtimeContainer *wf.WorkflowContainer) {

	router := gin.Default()
	setupCORS(router)

	v1 := router.Group("/api/v1")

	port := 8000

	// Register services / routes

	registerOverviewServices(v1, runtimeContainer)
	registerHealthServices(v1)
	registerEditorServices(v1)

	v1.GET("/test", func(c *gin.Context) {
		c.String(200, "SUC")
	})

	setupSwagger(router, port)

	// Setup completed

	fmt.Println("Starting server on port", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func setupSwagger(router *gin.Engine, port int) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func setupCORS(router *gin.Engine) {
	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		},
	))
}
