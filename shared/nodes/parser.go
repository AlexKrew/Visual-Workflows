package nodes

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"workflows/shared/shared_entities"
)

func ProcessParser(input *NodeInput, output *NodeOutput) error {

	inputValue, _ := input.ValueFor("input")
	outputType, _ := input.ValueFor("parse")

	var parsed any
	var parsingErr error

	// try parse
	switch inputValue.Value.(type) {
	case string:
		parsed, parsingErr = parseString(inputValue.Value.(string), outputType.Value.(string))
	case int:
		parsed, parsingErr = parseInt(inputValue.Value.(int), outputType.Value.(string))
	case float32:
		parsed, parsingErr = parseInt(int(inputValue.Value.(float32)), outputType.Value.(string))
	case float64:
		parsed, parsingErr = parseInt(int(inputValue.Value.(float64)), outputType.Value.(string))
	case bool:
		parsed, parsingErr = parseBool(inputValue.Value.(bool), outputType.Value.(string))
	}

	if parsingErr != nil {
		return parsingErr
	}

	var outputMessage shared_entities.WorkflowMessage
	switch outputType.Value.(string) {
	case "NUMBER":
		outputMessage = shared_entities.NumberMessage(parsed.(int))
	case "STRING":
		outputMessage = shared_entities.StringMessage(parsed.(string))
	case "BOOLEAN":
		outputMessage = shared_entities.BooleanMessage(parsed.(bool))
	}
	log.Println("PARSED VALUE:", parsed)
	output.Set("output", outputMessage)

	return nil
}

func parseString(s string, to string) (any, error) {
	var parsed any = nil
	var err error = nil

	switch to {
	case "NUMBER":
		parsed, err = stringToInt(s)
	case "BOOLEAN":
		parsed, err = stringToBool(s)
	}

	if parsed == nil {
		return parsed, errors.New("unknown type to parse to")
	}

	return parsed, err
}

func parseInt(i int, to string) (any, error) {
	log.Println("PARSE TO", to)
	var parsed any = nil
	var err error = nil

	switch to {
	case "STRING":
		parsed = intToString(i)
	case "BOOLEAN":
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
