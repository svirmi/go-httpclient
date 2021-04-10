package main

import (
	"fmt"

	"github.com/svirmi/go-httpclient/gohttp"
)

func main() {
	client := gohttp.HttpClient{}

	client.Get() // method invocation

	gohttp.Patch() // access to function

	fmt.Println(&client)
}
