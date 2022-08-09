package client

import "fmt"

func ProcessInject(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run Inject", input)
	output.Add("output", Message{Datatype: "STRING", Value: "http://google.com"})
}
