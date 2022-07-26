package goneis

import (
	"github.com/valyala/fastjson"
	"io"
	"net/http"
)

func newHTTPClient(key string) *HTTPClient {
	return &HTTPClient{
		key:    key,
		http:   &http.Client{},
		parser: &fastjson.Parser{},
	}
}

func (c HTTPClient) get(endpoint string, params *Parameters) (*fastjson.Value, error) {
	(*params)["TYPE"] = "json"
	request, err := http.NewRequest("GET", baseURL+endpoint+params.format(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	value, err := c.parser.ParseBytes(body)
	if err != nil {
		return nil, err
	}

	return value, nil
}
