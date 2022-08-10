package gcs

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"google.golang.org/api/iterator"
)

type GcsObjectManager struct {
	bucket string
}

func (objs GcsObjectManager) ListObjects(prefix string) ([]models.CloudStorageData, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.WithError(err).Error("Unable to connect to Google Cloud Storage")
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	it := client.Bucket(objs.bucket).Objects(ctx, &storage.Query{
		Prefix: prefix,
	})

	var result []models.CloudStorageData

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.WithError(err).Error("Error getting the list of objects")
			return nil, err
		}

		item := models.CloudStorageData{
			ObjectName: attrs.Name,
			ObjectData: "",
		}

		result = append(result, item)
	}

	return result, nil
}

func NewGcsObjectManager(bucket string) *GcsObjectManager {
	return &GcsObjectManager{
		bucket: bucket,
	}
}