package index

import (
	"context"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/sergeiten/medilastic/repository"
	log "github.com/sirupsen/logrus"
)

// Indexer ...
type Indexer interface {
	Do() error
}

// Indexer ...
type ElasticIndexer struct {
	ctx        context.Context
	client     *elastic.Client
	repository repository.Repository
	mapping    string
	indexName  string
	typeName   string
}

// NewIndexer ...
func NewIndexer(ctx context.Context, client *elastic.Client, r repository.Repository, mapping string, indexName string, typeName string) Indexer {
	return &ElasticIndexer{
		ctx:        ctx,
		client:     client,
		repository: r,
		mapping:    mapping,
		indexName:  indexName,
		typeName:   typeName,
	}
}

// Do stars indexing
func (i ElasticIndexer) Do() error {
	exists, err := i.client.IndexExists(i.indexName).Do(i.ctx)
	if err != nil {
		return err
	}

	if exists {
		_, err = i.client.DeleteIndex(i.indexName).Do(i.ctx)
		if err != nil {
			return err
		}
	}

	_, err = i.client.CreateIndex(i.indexName).BodyString(i.mapping).Do(i.ctx)
	if err != nil {
		return err
	}

	items, err := i.repository.Get()
	if err != nil {
		return err
	}

	for key, item := range items {
		p, err := i.client.Index().Index(i.indexName).Type(i.typeName).Id(strconv.Itoa(key)).BodyString(item).Do(i.ctx)
		if err != nil {
			log.WithError(err).Error("failed to insert item to index")
		}

		log.Infof("Indexed %s to index %s, type %s", p.Id, p.Index, p.Type)
	}

	return nil
}
