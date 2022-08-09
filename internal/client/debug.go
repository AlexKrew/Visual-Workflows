package client

import "fmt"

func ProcessDebug(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run Debug", input)

	for _, msg := range input.Values {
		ctx.Log(msg)
	}
	ctx.Log("Hello")
	output.Values = input.Values
}
