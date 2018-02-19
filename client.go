package medilastic

import (
	"context"

	"gopkg.in/olivere/elastic.v5"
)

// NewClient returns elastic client
func NewClient(ctx context.Context, url string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	_, _, err = client.Ping(url).Do(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
