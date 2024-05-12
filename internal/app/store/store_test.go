package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "admin:1234@/http_rest_api_test?charset=utf8"
	}
	os.Exit(m.Run())
}
