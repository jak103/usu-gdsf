package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"os"
	"testing"

	"github.com/fsouza/fake-gcs-server/fakestorage"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBucket(t *testing.T) {
	serverMock, err := fakestorage.NewServerWithOptions(fakestorage.Options{
		Scheme: "http",
	})
	if err != nil {
		log.WithError(err)
	}
	defer serverMock.Stop()

	url := serverMock.URL()

	os.Setenv("STORAGE_EMULATOR_HOST", url)
	instance := NewGcsBucketManager("test")

	bucket, err := instance.CreateBucket("TEST_BUCKET")
	if err != nil {
		assert.Fail(t, "Should be able to create a bucket that does not currently exist")
	}
	assert.NotNil(t, bucket)
	assert.NotEqual(t, uuid.Nil, bucket)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	// Testing if we can retrieve the created bucket
	bkt := client.Bucket("TEST_BUCKET")
	assert.NotNil(t, bkt)

	if err := os.Unsetenv("STORAGE_EMULATOR_HOST"); err != nil {
		log.WithError(err)
	}

}

func TestCreateNewBucketManagerWithoutServer(t *testing.T) {
	instance := NewGcsBucketManager("test")

	bucket, err := instance.CreateBucket("TEST_BUCKET")
	if assert.Error(t, err) {
		assert.ErrorContains(t, err, "dialing: google: could not find default credentials.")
	}
	assert.Equal(t, bucket, uuid.Nil)
}
