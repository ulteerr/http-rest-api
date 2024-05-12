package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseUrl = databaseUrl
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			query := fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))

			if _, err := s.db.Exec(query); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
