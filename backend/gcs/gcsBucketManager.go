package gcs

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
)

type GcsBucketManager struct {
	ProjectID string
}

func (creator GcsBucketManager) CreateBucket(bucketName string) (uuid.UUID, error) {
	// Connect GCS with the client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.WithError(err).Error("Error creating client")
		return uuid.Nil, err
	}
	defer client.Close()

	// Set our timeout to 30 seconds
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// Bucket Creation
	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, creator.ProjectID, nil); err != nil {
		log.WithError(err).Error("Unable to create bucket %v", bucket)
		return uuid.Nil, err
	}

	return uuid.New(), nil
}

func NewGcsBucketManager(projectID string) *GcsBucketManager {
	return &GcsBucketManager{
		ProjectID: projectID,
	}
}
