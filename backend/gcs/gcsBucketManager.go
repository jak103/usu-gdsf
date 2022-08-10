package gcs

import (
	"context"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"google.golang.org/api/option"
)

type GcsBucketManager struct {
	ProjectID string
	client *storage.Client
}

func (creator GcsBucketManager) CreateBucket(bucketName string) (uuid.UUID, error) {
	// Connect GCS with the client
	ctx := context.Background()

	if isDev, _ := strconv.ParseBool(os.Getenv("IS_DEV")); isDev {
		client, err := storage.NewClient(ctx, option.WithEndpoint(os.Getenv("STORAGE_EMULATOR_HOST") + "/storage/v1"))
		if err != nil {
			log.WithError(err).Error("Error creating client")
			return uuid.Nil, err
		}
		creator.client = client
	} else {
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.WithError(err).Error("Error creating client")
			return uuid.Nil, err
		}

		creator.client = client
	}

	defer creator.client.Close()

	// Set our timeout to 30 seconds
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// Bucket Creation
	bucket := creator.client.Bucket(bucketName)
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
