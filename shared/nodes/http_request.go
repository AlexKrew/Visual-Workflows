package nodes

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"workflows/shared/shared_entities"
)

func ProcessHttpRequest(input *NodeInput, output *NodeOutput) error {

	url, urlErr := input.ValueFor("url")
	method, methodErr := input.ValueFor("method")
	if urlErr != nil || methodErr != nil {
		fmt.Printf("http-request: input err: %s, %s", urlErr, methodErr)
		return errors.New("input error")
	}

	var response *http.Response
	var err error

	switch method.Value {
	case "GET":
		response, err = http.Get(url.Value.(string))
	}

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	output.Set("response-code", shared_entities.NumberMessage(response.StatusCode))
	output.Set("response", shared_entities.AnyMessage(string(dump)))

	return nil
}