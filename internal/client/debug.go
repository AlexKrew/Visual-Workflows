package client

func ProcessDebug(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	ctx.Log(input.Values["input"])
	output.Values = input.Values
}
