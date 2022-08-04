package gcs

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// downloadFile downloads an object to a file.
func downloadFile(bucket, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	// destFileName := "file.txt"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
					return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	/* f, err := os.Create(destFileName)
	if err != nil {
					return fmt.Errorf("os.Create: %v", err)
	}
	*/

	bla := client.Bucket(bucket)
	thing := bla.Object(object)
	rc, err := thing.NewReader(ctx)

	// rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
					return fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	/*
	if _, err := io.Copy(f, rc); err != nil {
					return fmt.Errorf("io.Copy: %v", err)
	}
	*/

	/*
	if err = f.Close(); err != nil {
					return fmt.Errorf("f.Close: %v", err)
	}
	*/

	// fmt.Fprintf(w, "Blob %v downloaded to local file %v\n", object, destFileName)

	return nil

}