package nodes

import (
	"errors"
	"fmt"
	"strconv"
	"workflows/shared/shared_entities"
)

func ProcessParse(input *NodeInput, output *NodeOutput) error {

	// TODO: Update identifiers
	inputValue, _ := input.ValueFor("input")
	outputType, _ := input.ValueFor("type")

	var parsed any
	var parsingErr error

	// try parse
	switch inputValue.Value.(type) {
	case string:
		parsed, parsingErr = parseString(inputValue.Value.(string), outputType.Value.(string))
	case int:
		parsed, parsingErr = parseInt(inputValue.Value.(int), outputType.Value.(string))
	case bool:
		parsed, parsingErr = parseBool(inputValue.Value.(bool), outputType.Value.(string))
	}

	if parsingErr != nil {
		return parsingErr
	}

	var outputMessage shared_entities.WorkflowMessage
	switch outputType.Value.(string) {
	case "int":
		outputMessage = shared_entities.NumberMessage(parsed.(int))
	case "string":
		outputMessage = shared_entities.StringMessage(parsed.(string))
	case "bool":
		outputMessage = shared_entities.BooleanMessage(parsed.(bool))
	}
	output.Set("parsed", outputMessage)

	return nil
}

func parseString(s string, to string) (any, error) {
	var parsed any = nil
	var err error = nil

	switch to {
	case "int":
		parsed, err = stringToInt(s)
	case "bool":
		parsed, err = stringToBool(s)
	}

	if parsed == nil {
		return parsed, errors.New("unknown type to parse to")
	}

	return parsed, err
}

func parseInt(i int, to string) (any, error) {
	var parsed any = nil
	var err error = nil

	switch to {
	case "string":
		parsed = intToString(i)
	case "bool":
		parsed, err = intToBool(i)
	}

	if parsed == nil {
		return parsed, errors.New("unknown type to parse to")
	}

	return parsed, err
}

func parseBool(b bool, to string) (any, error) {
	var parsed any = nil
	var err error = nil

	switch to {
	case "string":
		parsed = boolToString(b)
	case "int":
		parsed = boolToInt(b)
	}

	if parsed == nil {
		return parsed, errors.New("unknown type to parse to")
	}

	return parsed, err
}

// --- to string ---
func intToString(i int) string {
	return fmt.Sprint(i)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// --- to int ---
func stringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// --- to bool ---
func intToBool(i int) (bool, error) {
	if i == 0 {
		return false, nil
	}
	if i == 1 {
		return true, nil
	}

	return false, errors.New("could not convert int value to bool")
}

func stringToBool(s string) (bool, error) {
	if s == "true" || s == "1" {
		return true, nil
	}

	if s == "false" || s == "0" {
		return false, nil
	}

	return false, errors.New("could not convert string value to bool")
}
