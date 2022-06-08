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

	// Settings

	port := 8000

	// Register services / routes

	registerHealthServices(v1)
	registerOverviewServices(v1, runtimeContainer)
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
		},
	))
}

// func CORS() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
