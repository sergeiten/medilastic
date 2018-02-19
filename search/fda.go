package search

import (
	"context"
	"log"
	"reflect"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic"
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
func (s *fdaSearch) Search(query string, from int, size int) ([]map[string]string, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewMultiMatchQuery(query, "brand_name", "company_name", "device_description", "gmdn_pt_name", "gmdn_pt_definition", "product_code", "product_code_name").Fuzziness("AUTO").Operator("OR"))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	source, _ := searchQuery.Source()
	log.Printf("%+v", source)

	var result []map[string]string

	var ttyp medilastic.Fda
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(medilastic.Fda)

		d := map[string]string{
			"id":                 strconv.Itoa(t.ID),
			"brand_name":         t.BrandName,
			"company_name":       t.CompanyName,
			"device_description": t.DeviceDescription,
			"gmdn_pt_name":       t.GmdnPtName,
			"gmdn_pt_definition": t.GmdnPtDefinition,
			"product_code":       t.ProductCode,
			"product_code_name":  t.ProductCodeName,
		}

		result = append(result, d)
	}

	return result, nil
}
