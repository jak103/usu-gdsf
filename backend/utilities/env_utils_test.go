package utilities

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAsString(t *testing.T) {

	os.Setenv("SOME_ENV_VAR", "some env var value")
	defer os.Unsetenv("SOME_ENV_VAR")

	value := GetEnvAsString("SOME_ENV_VAR", "unused default")
	assert.Equal(t, "some env var value", value)

	value = GetEnvAsString("BAD_VAR", "used default")
	assert.Equal(t, "used default", value)
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("INT_VAR", "5")
	defer os.Unsetenv("INT_VAR")

	value := GetEnvAsInt("INT_VAR", 9)
	assert.Equal(t, 5, value)

	value = GetEnvAsInt("BAD_VAR", 9)
	assert.Equal(t, 9, value)

}
