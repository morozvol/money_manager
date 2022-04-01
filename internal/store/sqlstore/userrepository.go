package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {

	return r.store.db.QueryRowx(
		"INSERT INTO \"user\" (id,name) VALUES ($1,$2)",
		u.Id,
		u.Name,
	).Err()
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRowx(
		"SELECT id ,name FROM \"user\" WHERE id = $1",
		id,
	).StructScan(u); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
