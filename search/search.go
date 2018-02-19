package search

import (
	"context"

	"github.com/olivere/elastic"
)

// Searcher ...
type Searcher interface {
	Search(query string, from int, size int) ([]map[string]string, error)
}

func NewSearch(name string, ctx context.Context, client *elastic.Client) Searcher {
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
