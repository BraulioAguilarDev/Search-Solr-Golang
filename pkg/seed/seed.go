package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	c "search-engine/config"
	chttp "search-engine/pkg/occhttp"
	"time"
)

var URL = fmt.Sprintf("%s/jobs/update/json?commit=true", c.Config.SolrExternalHost)

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
