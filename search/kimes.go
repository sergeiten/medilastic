package search

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic"
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
func (s *kimesSearch) Search(query string, from int, size int) ([]map[string]string, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewMultiMatchQuery(query, "model", "country", "manufacture", "specification", "description", "category", "subcategory").Fuzziness("AUTO").Operator("AND"))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]string

	var ttyp medilastic.Kimes
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(medilastic.Kimes)

		d := map[string]string{
			"id":            strconv.Itoa(t.ID),
			"model":         t.Model,
			"country":       t.Country,
			"manufacture":   t.Manufacture,
			"specification": t.Specification,
			"description":   t.Description,
			"category":      t.Category,
			"subcategory":   t.Subcategory,
		}

		result = append(result, d)
	}

	return result, nil
}
