package main

import (
	"fmt"
	"io/ioutil"

	"github.com/svirmi/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	// headers := make(http.Header)
	// headers.Set("Authorization", "Bearer ABC-123")

	// client.SetHeaders(headers)

	return &client
}

func main() {

	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
