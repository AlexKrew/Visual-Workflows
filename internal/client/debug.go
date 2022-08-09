package client

import "fmt"

func ProcessDebug(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run Debug", input)

	ctx.Log(input.Values["input"].Value)
	output.Values = input.Values
}
