package store

import "github.com/ulteerr/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	_, err := r.store.db.Exec(
		"INSERT INTO users (email, encrypted_password) VALUES (?, ?)",
		u.Email,
		u.EncryptedPassword,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE email=?", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
