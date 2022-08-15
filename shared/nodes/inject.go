package nodes

func ProcessInject(input *NodeInput, output *NodeOutput) error {

	inputMsg, _ := input.ValueFor("input")
	output.Set("output", *inputMsg)

	return nil
}
