package search

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
)

// Searcher ...
type Searcher interface {
	Search(query string, from int, size int) ([]map[string]interface{}, error)
}

// NewSearch returns searcher interface for specific type
func NewSearch(ctx context.Context, name string, client *elastic.Client) Searcher {
	switch name {
	case "permit_status":
		return NewPermitStatusSearch(ctx, client)
	case "fda":
		return NewFdaSearch(ctx, client)
	case "kimes":
		return NewKimesSearch(ctx, client)
	case "medica":
		return NewMedicaSearch(ctx, client)
	case "pas":
		return NewPasSearch(ctx, client)
	case "pma":
		return NewPmaSearch(ctx, client)
	}
	return nil
}

func hitsResult(searchResult *elastic.SearchResult) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var t map[string]interface{}
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				return nil, err
			}

			for name, h := range hit.Highlight {
				fieldName := "highlight_" + name
				t[fieldName] = h
			}

			result = append(result, t)
		}
	}

	return result, nil
}

func searchQueryBuilder(query string, fields []string) *elastic.BoolQuery {
	return elastic.NewBoolQuery().Must(elastic.NewMultiMatchQuery(query, fields...).Operator("OR"))
}

func highlightBuilder() *elastic.Highlight {
	highlightFields := elastic.NewHighlighterField("*").PreTags("<em>").PostTags("</em>")
	return elastic.NewHighlight().NumOfFragments(0).Fields(highlightFields)
}
