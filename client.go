package medilastic

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

// NewClient returns elastic client
func NewClient(ctx context.Context) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	_, _, err = client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
