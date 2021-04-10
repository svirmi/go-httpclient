package main

import (
	"github.com/svirmi/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()

	gohttp.PublicFunc()

	client.Get()
}
