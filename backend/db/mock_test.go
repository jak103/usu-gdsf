package db

import "testing"

func TestMock(t *testing.T) {
	mock := Mock{}
	mock.Connect()
}
