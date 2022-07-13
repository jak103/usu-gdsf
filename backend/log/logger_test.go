package log

import (
	"testing"
)

type TestError struct {
	message string
}

func (te TestError) Error() string {
	return te.message
}

func TestLogger(t *testing.T) {
	t.Skip()
	Error("Error message")
	Error("WIth %s", "vars")
	err := TestError{message: "Test error"}
	WithError(err).Info("Some info message here")

}
