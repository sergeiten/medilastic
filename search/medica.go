package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// medicaSearch ...
type medicaSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewMedicaSearch ...
func NewMedicaSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &medicaSearch{
		ctx:       ctx,
		client:    client,
		indexName: "medica",
	}
}

func (s *medicaSearch) SetIndexName(name string) *medicaSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *medicaSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	fields := []string{
		"title", "description", "company_title", "company_description",
	}

	searchQuery := searchQueryBuilder(query, fields)
	highlight := highlightBuilder()

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).Highlight(highlight).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	result, err := hitsResult(searchResult)
	if err != nil {
		log.Printf("failed to convert hits to result: %v", err)
	}

	return result, nil
}
