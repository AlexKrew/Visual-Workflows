package nodes

import "visualWorkflows/internal/entities"

func Initialize(defaultDefinition *entities.Node, ownDefinition *entities.Node, ctx *entities.WorkflowContext) error {
	return nil
}

func Process(input *entities.Input, output *entities.Output, ctx *entities.WorkflowContext) error {
	ctx.Logger.Info("Execute Inject Node")
	return nil
}

func BeforeShutdown(ctx *entities.WorkflowContext) error {
	return nil
}
