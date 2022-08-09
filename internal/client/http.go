package client

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func ProcessHttp(input *NodeInput, output *NodeOutput, ctx *NodeContext) {
	fmt.Println("Run HTTP", input)

	url := input.Values["url"]
	method := input.Values["method"]

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

	output.Add("response-code", Message{Datatype: "NUMBER", Value: response.StatusCode})
	output.Add("response", Message{Datatype: "ANY", Value: string(dump)})
}
