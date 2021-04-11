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

	fullHeaders := c.setRequestHeaders(headers)
	request.Header = fullHeaders

	return client.Do(request)
}

func (c *httpclient) setRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// set default headers from the HTTP client instance to the Request
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// set custom headers from the current request to the Request (over-write defaults)
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}
