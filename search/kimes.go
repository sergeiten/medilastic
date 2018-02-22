package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// kimesSearch ...
type kimesSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewKimesSearch ...
func NewKimesSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &kimesSearch{
		ctx:       ctx,
		client:    client,
		indexName: "kimes",
	}
}

func (s *kimesSearch) SetIndexName(name string) *kimesSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *kimesSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	fields := []string{
		"model", "country", "manufacture", "specification", "description", "category", "subcategory",
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
