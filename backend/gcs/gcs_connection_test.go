package gcs

import (
	"os"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	os.Setenv("STORAGE_EMULATOR_HOST", "http://localhost:4443/storage/v1/b")
	err := downloadFile("breakout", "index.html")
	if err != nil {
		t.Fatalf("Failed to download file.")
	}

	if err := os.Unsetenv("STORAGE_EMULATOR_HOST"); err != nil {
		t.Fatalf("STORAGE_EMULATOR_HOST could not be unset.")
	}
}
