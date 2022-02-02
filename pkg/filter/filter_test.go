package filter_test

import (
	"search-engine/pkg/filter"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParserQuery  a simple test about queryParams that we pass by query string
func TestParserQuery(t *testing.T) {
	allowed := []string{
		"title",
		"text",
		"salary",
		"locality",
		"category",
	}

	tests := []struct {
		input          string
		expectedString string
		expectedError  string
	}{
		{
			"",
			"*",
			"",
		},
		{
			"locality=remoto&category=dev",
			"locality:remoto+category:dev+",
			"",
		},
		{
			"test=go",
			"",
			"key test invalid",
		},
	}

	for _, test := range tests {
		result, err := filter.ParserQuery(test.input, allowed)

		assert.Equal(t, test.expectedString, result)
		if err != nil {
			assert.Equal(t, test.expectedError, err.Error())
		}
	}
}
