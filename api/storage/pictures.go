package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
)

var (
	STORAGE_BUCKET_NAME = os.Getenv("STORAGE_BUCKET_NAME")
)

func UploadPicture(pictureBytes []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	namePicture := fmt.Sprintf("%s.png", uuid.New())
	writterClient := clientStorage.Bucket(STORAGE_BUCKET_NAME).Object(namePicture).NewWriter(ctx)
	writterClient.ChunkSize = 0

	if _, err := io.Copy(writterClient, bytes.NewBuffer(pictureBytes)); err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	if err := writterClient.Close(); err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	return fmt.Sprintf("/%s/%s", STORAGE_BUCKET_NAME, namePicture), nil
}

func DeletePicture(pictureUrl string) error {
	if strings.Contains(pictureUrl, STORAGE_BUCKET_NAME) {
		pictureUrl = strings.Replace(pictureUrl, fmt.Sprintf("/%s/", STORAGE_BUCKET_NAME), "", -1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	err := clientStorage.Bucket(STORAGE_BUCKET_NAME).Object(pictureUrl).Delete(ctx)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}
