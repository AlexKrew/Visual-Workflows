package nodes

import (
	"fmt"
	"log"
)

func ProcessDebug(input *NodeInput, output *NodeOutput) error {

	logMessage, err := input.ValueFor("input")
	if err != nil {
		log.Printf("debug failed: %s", err)
		return err
	}

	output.Log(fmt.Sprintf("DEBUG: %v", logMessage.Value))

	output.Set("output", *logMessage)

	return nil
}
