package search

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

// PermitStatusSearch ...
type permitStatusSearch struct {
	ctx       context.Context
	client    *elastic.Client
	indexName string
}

// NewPermitStatusSearch ...
func NewPermitStatusSearch(ctx context.Context, client *elastic.Client) Searcher {
	return &permitStatusSearch{
		ctx:       ctx,
		client:    client,
		indexName: "permit_status",
	}
}

func (s *permitStatusSearch) SetIndexName(name string) *permitStatusSearch {
	s.indexName = name
	return s
}

// Search ...
func (s *permitStatusSearch) Search(query string, from int, size int) ([]map[string]interface{}, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewQueryStringQuery(query).DefaultField("*").AnalyzeWildcard(true))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	// var result []map[string]string

	// ID             int    `json:"id"`
	// Product        string `json:"product"`
	// Entrps         string `json:"entrps"`
	// PrductPrmisnNo string `json:"prduct_prmisn_no"`
	// MeaClassNo     string `json:"mea_class_no"`
	// TypeName       string `json:"type_name"`
	// UsePurps       string `json:"use_purps"`
	// var ttyp medilastic.PermitStatus
	// for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	// 	t := item.(medilastic.PermitStatus)

	// 	d := map[string]string{
	// 		"id":               strconv.Itoa(t.ID),
	// 		"product":          t.Prduct,
	// 		"entrps":           t.Entrps,
	// 		"prduct_prmisn_no": t.PrductPrmisnNo,
	// 		"mea_class_no":     t.MeaClassNo,
	// 		"type_name":        t.TypeName,
	// 		"use_purps":        t.UsePurps,
	// 	}

	// 	result = append(result, d)
	// }

	result, err := hitsResult(searchResult)
	if err != nil {
		log.Printf("failed to convert hits to result: %v", err)
	}

	return result, nil
}
