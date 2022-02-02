package solr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search-engine/models"
	chttp "search-engine/pkg/occhttp"
	"strings"
	"time"

	"github.com/golang/glog"
)

type SearchRepository struct {
	URL   string
	Debug bool
}

func NewSearchRepository(url string, debug bool) *SearchRepository {
	return &SearchRepository{
		URL:   url,
		Debug: debug,
	}
}

// Find func executes a request to external service with http client
func (s *SearchRepository) Find(query string) (*models.SolrResponse, error) {
	chttp.SetDebugMode(s.Debug)

	// Parse URL with qs
	URL := strings.TrimSuffix(fmt.Sprintf("%s?q=%s", s.URL, query), "+")

	// Config options
	opts := &chttp.Options{
		Timeout:  5 * time.Second,
		Endpoint: URL,
		Method:   http.MethodGet,
	}

	client := chttp.NewClient()

	result, err := client.MakeRequest(opts)
	if err != nil {
		glog.V(5).Info(fmt.Sprintf("Make request error: %v", err))
		return nil, err
	}

	var searchResult *models.SolrResponse

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		glog.V(5).Info(fmt.Sprintf("ReadAll error: %v", err))
		return nil, err
	}

	if err := json.Unmarshal([]byte(body), &searchResult); err != nil {
		glog.V(5).Info(fmt.Sprintf("Unmarshal error: %v", err))
		return searchResult, nil
	}

	return searchResult, nil
}
