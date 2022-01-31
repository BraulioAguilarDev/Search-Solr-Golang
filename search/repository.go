package search

import "search-engine/models"

type Repository interface {
	Find(query string) (*models.SolrResponse, error)
}
