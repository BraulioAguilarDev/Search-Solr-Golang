package search

import (
	"search-engine/search/presenter"
)

type UseCase interface {
	Find(query string) (presenter.SearchResponse, error)
}
