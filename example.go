package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/svirmi/go-httpclient/gohttp"
)

var (
	githubHttpClient = gohttp.New()
)

func main() {

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := githubHttpClient.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
