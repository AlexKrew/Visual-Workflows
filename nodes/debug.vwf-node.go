package nodes

import (
	"fmt"
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/node"
)

// func Initialize(defaultDefinition *entities.Node, ownDefinition *entities.Node, ctx *entities.WorkflowContext) error {
// 	return nil
// }

func ProcessDebug(input entities.Input, output *entities.Output, ctx *node.WorkflowContextProxy) error {

	msgToLog := input.GetMessage("input")

	ctx.Log(fmt.Sprintf("%v", msgToLog)) // convert any interface to string
	fmt.Println("Debug:", msgToLog)

	output.Add("output", msgToLog)
	return nil
}

// func BeforeShutdown(ctx *entities.WorkflowContext) error {
// 	return nil
// }
