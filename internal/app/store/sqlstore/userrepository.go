package sqlstore

import (
	"database/sql"

	"github.com/ulteerr/http-rest-api/internal/app/model"
	"github.com/ulteerr/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	_, err := r.store.db.Exec(
		"INSERT INTO users (email, encrypted_password) VALUES (?, ?)",
		u.Email,
		u.EncryptedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE email=?", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
