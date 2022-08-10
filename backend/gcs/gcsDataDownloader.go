package gcs

import (
	"context"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jak103/usu-gdsf/log"
)

type GcsDataDownloader struct {
}

func (download GcsDataDownloader) DownloadData(bucket, object string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.WithError(err).Error("Unable to connect to GCS")
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second * 50)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		log.WithError(err).Error("Unable to read from Bucket")
		return nil, err
	}
	defer rc.Close()

	if data, err := ioutil.ReadAll(rc); err != nil {
		log.WithError(err).Error("An error occurred while reading the data into memory")
		return nil, err
	} else {
		return data, nil
	}
}

func NewGcsDataDownloader() *GcsDataDownloader {
	return &GcsDataDownloader{
	}
}