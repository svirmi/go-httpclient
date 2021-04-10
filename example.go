package main

import (
	"fmt"

	"github.com/svirmi/go-httpclient/gohttp"
)

func main() {
	client := gohttp.HttpClient{}
	fmt.Println(&client)
}
