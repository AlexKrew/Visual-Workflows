package server

import (
	"net/http"

	container "visualWorkflows/internal/container"

	"github.com/gin-gonic/gin"
)

var wfContainer *container.WorkflowContainer

func registerOverviewServices(rg *gin.RouterGroup, container *container.WorkflowContainer) {
	wfContainer = container

	overviewServices := rg.Group("/workflows")

	overviewServices.GET("", getWorkflows)
}

func getWorkflows(c *gin.Context) {

	wfInfos, err := wfContainer.GetAvailableWorkflows()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, wfInfos)
}
