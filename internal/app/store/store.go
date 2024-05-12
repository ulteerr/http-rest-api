package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}
