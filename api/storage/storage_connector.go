package storage

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/getsentry/sentry-go"
	"google.golang.org/api/option"
)

var (
	clientStorage   *storage.Client
	onceConnectToDB sync.Once
)

func init() {
	onceConnectToDB.Do(
		func() {
			bucketClient, err := connectToBucket()
			if err != nil {
				sentry.CaptureException(err)
			}
			clientStorage = bucketClient
		},
	)
}

func connectToBucket() (*storage.Client, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile("./key.json"))
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	defer client.Close()

	return client, nil
}
