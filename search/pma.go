package search

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic"
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
func (s *pmaSearch) Search(query string, from int, size int) ([]map[string]string, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewMultiMatchQuery(query, "applicant", "generic_name", "trade_name").Fuzziness("AUTO").Operator("AND"))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]string

	var ttyp medilastic.Pma
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(medilastic.Pma)

		d := map[string]string{
			"id":           strconv.Itoa(t.ID),
			"applicant":    t.Applicant,
			"generic_name": t.GenericName,
			"trade_name":   t.TradeName,
		}

		result = append(result, d)
	}

	return result, nil
}
