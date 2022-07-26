package nodes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/node"
)

// func Initialize(defaultDefinition *entities.Node, ownDefinition *entities.Node, ctx *entities.WorkflowContext) error {
// 	return nil
// }

func ProcessHttpRequest(input entities.Input, output *entities.Output, ctx *node.WorkflowContextProxy) error {

	url := input.GetMessage("url")

	fmt.Println("Http: Get", url.Value)

	response, err := http.Get(url.Value.(string))
	if err != nil {
		fmt.Println("Failed to fetch url", err)
		return err
	}

	statusCode := response.StatusCode
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println("Failed to dump response", err)
		return err
	}

	output.Add("code", entities.NumberMessage(statusCode))
	output.Add("data", entities.StringMessage(string(dump)))

	return nil
}

// func BeforeShutdown(ctx *entities.WorkflowContext) error {
// 	return nil
// }
