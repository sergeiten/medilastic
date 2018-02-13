package main

import (
	"context"
	"flag"

	"github.com/sergeiten/medilastic"
	"github.com/sergeiten/medilastic/print"
	"github.com/sergeiten/medilastic/search"

	log "github.com/sirupsen/logrus"
)

var query string
var indexName string
var from int
var size int

func init() {
	flag.StringVar(&query, "query", "", "Search Query")
	flag.StringVar(&indexName, "index", "", "Index Name")
	flag.IntVar(&from, "from", 0, "Start from position")
	flag.IntVar(&size, "size", 10, "Limit of returned items")
	flag.Parse()

	if query == "" {
		log.Fatal("search query cannot be empty")
	}
}

func main() {
	ctx := context.Background()

	client, err := medilastic.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("failed to get elastic client")
	}

	search := search.NewSearch(indexName, ctx, client)

	result, err := search.Search(query, from, size)
	if err != nil {
		log.WithError(err).Fatal("failed to search")
	}

	var fields []string

	switch indexName {
	case "permit_status":
		fields = []string{"ID", "Product", "Ent", "PerNo", "Class NO", "Type", "Use"}
	case "fda":
		fields = []string{"ID", "BrandName", "CompanyName", "DeviceDescription", "GmdnPtName", "GmdnPtDefinition", "ProductCode", "ProductCodeName"}
	}

	print := print.New(fields, result)
	print.Print()
}
