package webserver

import (
	"fmt"
	"log"
	"net/http"

	"workflows/internal/dashboard"
	"workflows/internal/processors/workflow_processor"
	"workflows/internal/workflows"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/reactivex/rxgo/v2"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

var WFHelper WorkflowHelper

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @title Visual Workflow API
// @version 0.1
// @description description

// @host localhost:8000
// @BasePath /api/v1
func StartServer(eventStream *workflows.EventStream, workflowProcessor *workflow_processor.WorkflowProcessor) {

	WFHelper = ConstructWorkflowHelper(workflowProcessor)

	router := gin.Default()
	setupCORS(router)

	v1 := router.Group("/api/v1")

	port := 8000

	builderEvents := make(chan any, 100)
	dashboardEvents := make(chan any, 100)
	go registerEventsHandler(eventStream.EventsObservable, builderEvents, dashboardEvents)

	go setupBuilderWebsocket(router, builderEvents)
	go setupDashboardWebsocket(router, dashboardEvents)

	// Register services / routes

	registerOverviewServices(v1)
	registerWorkflowServices(v1)
	registerHealthServices(v1)
	registerEditorServices(v1)
	registerDashboardServices(v1)

	v1.GET("/test", func(c *gin.Context) {
		c.String(200, "Online")
	})

	// setupSwagger(router, port)

	// Setup completed

	fmt.Println("Starting server on port", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// func setupSwagger(router *gin.Engine, port int) {
// 	docs.SwaggerInfo.BasePath = "/api/v1"
// 	ginSwagger.WrapHandler(swaggerFiles.Handler,
// 		ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)),
// 		ginSwagger.DefaultModelsExpandDepth(-1))

// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
// }

func setupCORS(router *gin.Engine) {
	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"*"},
		},
	))
}

func setupBuilderWebsocket(router *gin.Engine, events chan any) {
	router.GET("/workflow/websocket", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}
		defer ws.Close()

		for {
			ev := <-events
			err = ws.WriteJSON(ev)
			if err != nil {
				panic(err)
			}
		}

	})
}

func setupDashboardWebsocket(router *gin.Engine, events chan any) {
	router.GET("/dashboard/websocket", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}
		defer ws.Close()

		for {
			ev := <-events
			fmt.Println("PUB DB EVENT")

			err = ws.WriteJSON(ev)
			if err != nil {
				panic(err)
			}
		}

	})
}

func registerEventsHandler(observable *rxgo.Observable, builderEvents chan any, dashboardEvents chan any) {
	(*observable).ForEach(func(i interface{}) {
		event := i.(workflows.WorkflowEvent)

		switch event.Type {
		case workflows.DebugEvent:
			body := event.Body.(workflows.DebugEventBody)

			message := make(map[string]any)
			message["id"] = event.ID
			message["timestamp"] = event.CreatedAt
			message["message"] = body.Value
			// TODO: Add workflow id
			log.Println("DEBUG EVENT")

			builderEvents <- message

		case workflows.DashboardValueChanged:
			body := event.Body.(workflows.DashboardValueChangedEventBody)

			data := make(map[string]any)
			data["id"] = body.ElementID
			data["field"] = body.Field
			data["value"] = body.Value

			message := make(map[string]any)
			message["workflow_id"] = body.WorkflowID
			message["type"] = "field_updated"
			message["data"] = data

			dashboardEvents <- message

		case workflows.WorkflowReady:

			body := event.Body.(workflows.WorkflowReadyEventBody)

			workflow, exists := WFHelper.WorkflowById(body.WorkflowID)
			if !exists {
				return
			}
			config, exists := dashboard.ConfigFromWorkflow(*workflow)
			if !exists {
				return
			}

			message := make(map[string]any)
			message["workflow_id"] = body.WorkflowID
			message["type"] = "rebuild_ui"
			message["data"] = config

			fmt.Println("ADD REBUILD UI")
			dashboardEvents <- message
		}

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}
