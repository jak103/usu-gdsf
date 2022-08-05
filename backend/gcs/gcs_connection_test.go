package gcs

import (
	"os"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	os.Setenv("STORAGE_EMULATOR_HOST", "http://localhost:4443")
	err := downloadFile("breakout", "game/index.html")
	if err != nil {
		print("bla")
	}

	if err := os.Unsetenv("STORAGE_EMULATOR_HOST"); err != nil {
		print("bla")
	}
}
