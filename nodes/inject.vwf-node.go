package nodes

import (
	"fmt"
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/node"
)

// func Initialize(defaultDefinition *entities.Node, ownDefinition *entities.Node, ctx *entities.WorkflowContext) error {
// 	return nil
// }

func ProcessInject(input entities.Input, output *entities.Output, ctx *node.WorkflowContextProxy) error {

	msg := input.GetMessage("input")
	fmt.Println("Inject", msg)

	return nil
}

// func BeforeShutdown(ctx *entities.WorkflowContext) error {
// 	return nil
// }
