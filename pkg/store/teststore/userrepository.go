package teststore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
