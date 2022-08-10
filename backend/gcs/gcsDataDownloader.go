package gcs

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jak103/usu-gdsf/log"
	"google.golang.org/api/option"
)

type GcsDataDownloader struct {
	client *storage.Client
}

func (download GcsDataDownloader) DownloadData(bucket, object string) ([]byte, error) {
	ctx := context.Background()
	if isDev, _ := strconv.ParseBool(os.Getenv("IS_DEV")); isDev {
		client, err := storage.NewClient(ctx, option.WithEndpoint(os.Getenv("STORAGE_EMULATOR_HOST") + "/storage/v1"))
		if err != nil {
			log.WithError(err).Error("Error creating client")
			return nil, err
		}

		download.client = client
	} else {
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.WithError(err).Error("Error creating client")
			return nil, err
		}

		download.client = client
	}

	defer download.client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second * 50)
	defer cancel()

	rc, err := download.client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		log.Debug("Object: %v", object)
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