package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// pmaSearch ...
type pmaSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewPmaSearch ...
func NewPmaSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &pmaSearch{
		ctx:       ctx,
		client:    client,
		indexName: "pma",
	}
}

func (s *pmaSearch) SetIndexName(name string) *pmaSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *pmaSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	fields := []string{
		"applicant", "generic_name", "trade_name",
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
