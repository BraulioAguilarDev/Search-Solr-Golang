package usecase

import (
	http "search-engine/pkg/occhttp"
	"search-engine/search"
	"search-engine/search/presenter"

	"github.com/golang/glog"
)

type SearchUseCase struct {
	SearchRepo search.Repository
}

func NewUseCase(repo search.Repository) *SearchUseCase {
	return &SearchUseCase{
		SearchRepo: repo,
	}
}

// Find function returns a list of rows filtered by query string
func (s SearchUseCase) Find(query string) (presenter.SearchResponse, error) {

	rows, err := s.SearchRepo.Find(query)
	if err != nil {
		glog.Error(err)

		return presenter.SearchResponse{}, err
	}

	// Map result to friendly struct
	data := presenter.MapDocsToDataResponse(rows)

	response := presenter.SearchResponse{
		PayloadResponseEnvelope: http.PayloadResponseEnvelope{
			Success: true,
		},
		Data: data,
	}

	return response, nil
}
