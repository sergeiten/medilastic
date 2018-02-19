package search

import (
	"context"
	"reflect"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic"
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
func (s *pasSearch) Search(query string, from int, size int) ([]map[string]string, error) {
	searchQuery := elastic.NewBoolQuery()
	searchQuery.Must(elastic.NewMultiMatchQuery(query, "application_name", "device_name", "medical_speciality", "study_name", "study_design_description").Fuzziness("AUTO").Operator("AND"))

	searchResult, err := s.client.Search().Index(s.indexName).Query(searchQuery).From(from).Size(size).Do(s.ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]string

	var ttyp medilastic.Pas
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(medilastic.Pas)

		d := map[string]string{
			"id":                       strconv.Itoa(t.ID),
			"application_name":         t.ApplicationName,
			"device_name":              t.DeviceName,
			"medical_speciality":       t.MedicalSpeciality,
			"study_name":               t.StudyName,
			"study_design_description": t.StudyDesignDescription,
		}

		result = append(result, d)
	}

	return result, nil
}
