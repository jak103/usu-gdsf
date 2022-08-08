package gcs

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
)

type CloudDataUploader struct {
	projectID string
}

func (up CloudDataUploader) UploadFile(bucket, object string, data []byte) (uuid.UUID, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.WithError(err).Error("An error occurred while connecting to GCS")
		return uuid.Nil, err
	}
	defer client.Close()

	buf := bytes.NewBuffer(data)

	ctx, cancel := context.WithTimeout(ctx, time.Second * 50)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0

	if _, err = io.Copy(wc, buf); err != nil {
		log.WithError(err).Error("An error occurred while writing to GCS")
		return uuid.Nil, err
	}

	if err := wc.Close(); err != nil {
		log.WithError(err).Error("An error occurred while closing the stream.")
		return uuid.Nil, err
	}

	log.Info("%v uploaded to bucket: %v", object, bucket)
	return uuid.New(), nil
}

func NewCloudDataUploader(projectID string) *CloudDataUploader {
	return &CloudDataUploader{
		projectID: projectID,
	}
}