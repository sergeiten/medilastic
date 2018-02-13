package search

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
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
	}
	return nil
}
