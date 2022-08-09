package client

import "fmt"

func ProcessHttp(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run HTTP", input)
}
