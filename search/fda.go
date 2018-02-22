package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// FdaSearch ...
type fdaSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewFdaSearch ...
func NewFdaSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &fdaSearch{
		ctx:       ctx,
		client:    client,
		indexName: "fda",
	}
}

func (s *fdaSearch) SetIndexName(name string) *fdaSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *fdaSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	fields := []string{
		"brand_name",
		"company_name",
		"device_description",
		"gmdn_pt_name",
		"gmdn_pt_definition",
		"product_code",
		"product_code_name",
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
