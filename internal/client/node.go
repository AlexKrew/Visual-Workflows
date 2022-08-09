package client

type NodeFunc func(input any, output any, ctx any)

type NodeInput struct {
	Values map[string]Message
}

func NewNodeInput(input map[string]Message) NodeInput {
	return NodeInput{
		Values: input,
	}
}

type NodeOutput struct {
	Values map[string]Message
}

func NewNodeOutput() NodeOutput {
	return NodeOutput{
		Values: make(map[string]Message),
	}
}

func (output *NodeOutput) Add(key string, message Message) {
	output.Values[key] = message
}

type NodeContext struct {
	Logs []any
}

func NewNodeContext() NodeContext {
	return NodeContext{
		Logs: []any{},
	}
}

func (ctx *NodeContext) Log(value any) {
	ctx.Logs = append(ctx.Logs, value)
}
