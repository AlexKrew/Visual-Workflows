package nodes

import (
	"fmt"
	"log"
	"strconv"
	"workflows/shared/shared_entities"
)

func ProcessIf(input *NodeInput, output *NodeOutput) error {

	checkValue, _ := input.ValueFor("check")
	operatorValue, _ := input.ValueFor("operator")
	compareValue, _ := input.ValueFor("value")

	var shouldTrigger bool

	switch checkValue.Value.(type) {
	case string:
		shouldTrigger = ifString(checkValue.Value.(string), operatorValue.Value.(string), *compareValue)
	case float64:
		shouldTrigger = ifNumber(checkValue.Value.(float64), operatorValue.Value.(string), *compareValue)
	default:
		fmt.Printf("datatype `%s` is not supported", checkValue.DataType)
	}

	if !shouldTrigger {
		output.SetTrigger(false)
	}

	return nil
}

func ifString(v1 string, op string, v2 shared_entities.WorkflowMessage) bool {
	v2Value := fmt.Sprintf("%v", v2.Value)

	switch op {
	case ">=":
		return len(v1) >= len(v2Value)
	case ">":
		return len(v1) > len(v2Value)
	case "==":
		return v1 == v2Value
	case "!=":
		return v1 != v2Value
	case "<=":
		return len(v1) <= len(v2Value)
	case "<":
		return len(v1) < len(v2Value)
	}

	fmt.Printf("Operator `%s` is not supported for type STRING", op)

	return false
}

func ifNumber(v1 float64, op string, v2 shared_entities.WorkflowMessage) bool {
	var v2Value float64

	switch v2.DataType {
	case "STRING":
		v2Int, err := strconv.Atoi(v2.Value.(string))
		if err != nil {
			log.Printf("Failed to convert %s to float", v2.Value)
			return false
		}
		v2Value = float64(v2Int)

	case "NUMBER":
		v2Value = v2.Value.(float64)

	default:
		fmt.Printf("compare datatype `%s` is not supported for type number", v2.DataType)
		return false
	}

	switch op {
	case ">=":
		return v1 >= v2Value
	case ">":
		return v1 > v2Value
	case "==":
		return v1 == v2Value
	case "!=":
		return v1 != v2Value
	case "<=":
		return v1 <= v2Value
	case "<":
		return v1 <= v2Value
	}

	fmt.Printf("Operator `%s` is not supported for type NUMBER", op)

	return false
}
