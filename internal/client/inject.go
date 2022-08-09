package client

func ProcessInject(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	output.Add("output", Message{Datatype: "STRING", Value: "http://google.com"})
}
