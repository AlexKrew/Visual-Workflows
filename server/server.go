package server

import (
	"fmt"
	"net/http"

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

	events := make(chan any)
	go registerEventsHandler(eventStream.EventsObservable, &events)

	go setupWebsocket(router, events)

	// Register services / routes

	registerOverviewServices(v1)
	registerWorkflowServices(v1)
	registerHealthServices(v1)
	registerEditorServices(v1)

	v1.GET("/test", func(c *gin.Context) {
		c.String(200, "SUC")
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

func setupWebsocket(router *gin.Engine, events chan any) {
	router.GET("/websocket", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}
		defer ws.Close()

		for ev := range events {
			fmt.Println("PUSH WS EVENT")
			err = ws.WriteJSON(ev)
			if err != nil {
				panic(err)
			}
		}

	})
}

func registerEventsHandler(observable *rxgo.Observable, debugEvents *chan any) {
	(*observable).ForEach(func(i interface{}) {
		event := i.(workflows.WorkflowEvent)

		fmt.Println("SERVER EVENT", event.Type, workflows.DebugEvent)

		switch event.Type {
		case workflows.DebugEvent:
			fmt.Println("DEBUG EVENT IN SERVER", event)

			body := event.Body.(workflows.DebugEventBody)

			message := make(map[string]any)
			message["id"] = event.ID
			message["timestamp"] = event.CreatedAt
			message["message"] = body.Value

			(*debugEvents) <- message

		}

	}, func(err error) {
		fmt.Println("Error", err)

	}, func() {})
}
