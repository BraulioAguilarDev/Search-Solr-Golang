package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	chttp "search-engine/pkg/occhttp"
	"time"
)

var URL = "http://localhost:8983/solr/jobs/update/json?commit=true"

func seed() (*http.Response, error) {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, _ := io.ReadAll(jsonFile)

	options := &chttp.Options{
		Timeout:  3 * time.Second,
		Endpoint: URL,
		Method:   http.MethodPost,
		Body:     data,
	}

	client := chttp.NewClient()
	res, err := client.MakeRequest(options)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	result, err := seed()
	if err != nil {
		fmt.Printf("seed error: %v\n", err)
	} else {
		fmt.Printf("successfull action -- code %d\n", result.StatusCode)
	}
}
