package gohttp

import "net/http"

type httpclient struct{}

// only New() exported and publically available outside this package
func New() httpclient {
	client := &httpclient{}
	return *client
}

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

// any struct implemented these methods will be concidered HttpClient

// implementations of interface
func (c *httpclient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpclient) Post() {}

func (c *httpclient) Put() {}

func (c *httpclient) Patch() {}

func (c *httpclient) Delete() {}

func PublicFunc() {

}
