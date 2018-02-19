package search

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic"
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
func (s *medicaSearch) Search(query string, from int, size int) ([]map[string]string, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewMultiMatchQuery(query, "title", "description", "company_title", "company_description").Fuzziness("AUTO").Operator("AND"))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]string

	var ttyp medilastic.Medica
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(medilastic.Medica)

		d := map[string]string{
			"id":                  strconv.Itoa(t.ID),
			"title":               t.Title,
			"description":         t.Description,
			"company_title":       t.CompanyTitle,
			"company_description": t.CompanyDescription,
		}

		result = append(result, d)
	}

	return result, nil
}
