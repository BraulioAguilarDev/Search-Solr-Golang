package presenter

import (
	"search-engine/models"
	chttp "search-engine/pkg/occhttp"
)

type Data struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Salary   int    `json:"salary"`
	Locality string `json:"locality"`
	Category string `json:"category"`
}

type SearchResponse struct {
	chttp.PayloadResponseEnvelope
	Data []Data `json:"data"`
}

func MapDocsToDataResponse(docs *models.SolrResponse) (data []Data) {

	for _, d := range docs.Response.Docs {
		data = append(data, Data{
			Title:    d.Title[0],
			Text:     d.Text[0],
			Salary:   d.Salary[0],
			Locality: d.Locality[0],
			Category: d.Category[0],
		})
	}

	return data
}
