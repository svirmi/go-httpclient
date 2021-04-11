package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpclient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("Unable to create new request")
	}

	return client.Do(request)
}
