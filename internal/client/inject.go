package client

import "fmt"

func ProcessInject(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run Inject", input)
}
