package filter

import (
	"fmt"
	"strings"
)

// Simulates InArray behavior at PHP -- checks if exist a value
func Exists(items []string, val string) bool {
	for _, item := range items {
		if item == val {
			return true
		}
	}

	return false
}

// ParserQuery parse our params (title, category, text, salary, locality)
// to solr service
func ParserQuery(queryStringURL string, fields []string) (string, error) {

	var qFiler string

	// example: ?category=data&title=base de datos
	if len(queryStringURL) == 0 {
		return "*", nil
	}

	// to split when exists '&'
	blocks := strings.Split(queryStringURL, "&")

	for _, v := range blocks {
		// get key & value
		filter := strings.Split(v, "=")
		key := filter[0]
		value := filter[1]

		// Validate if this key is allowed for external service
		if !Exists(fields, key) {
			return "", fmt.Errorf("key %s invalid", key)
		}

		qFiler = qFiler + fmt.Sprintf("%s:%s+", key, value)
	}

	return qFiler, nil
}
