package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// pasSearch ...
type pasSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewPasSearch ...
func NewPasSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &pasSearch{
		ctx:       ctx,
		client:    client,
		indexName: "pas",
	}
}

func (s *pasSearch) SetIndexName(name string) *pasSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *pasSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	fields := []string{
		"application_name", "device_name", "medical_speciality", "study_name", "study_design_description",
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
