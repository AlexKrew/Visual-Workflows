package nodes

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"workflows/shared/shared_entities"
)

func ProcessHttpRequest(input *NodeInput, output *NodeOutput) error {

	url, urlErr := input.ValueFor("url")
	if urlErr != nil {
		fmt.Printf("http-request: input err: %s", urlErr)
		return errors.New("input error")
	}

	var response *http.Response
	var err error

	urlValue := url.Value.(string)
	if !strings.HasPrefix(urlValue, "http") || !strings.HasPrefix(urlValue, "https") {
		urlValue = fmt.Sprintf("http://%s", urlValue)
	}

	response, err = http.Get(urlValue)

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	output.Set("response_code", shared_entities.NumberMessage(response.StatusCode))
	output.Set("response", shared_entities.AnyMessage(string(dump)))

	return nil
}
