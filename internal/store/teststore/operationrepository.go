package teststore

import (
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

type OperationRepository struct {
	store      *Store
	operations map[int]*model.Operation
}

// Create ...
func (r *OperationRepository) Create(o *model.Operation) error {

	o.Id = int64(len(r.operations) + 1)
	r.operations[int(o.Id)] = o

	return nil
}

// Find ...
func (r *OperationRepository) Find(id int) (*model.Operation, error) {
	o, ok := r.operations[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return o, nil
}
