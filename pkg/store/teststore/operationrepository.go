package teststore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

type OperationRepository struct {
	store      *Store
	operations map[int]*model.Operation
}

// Create ...
func (r *OperationRepository) Create(o ...*model.Operation) error {
	for _, oper := range o {
		oper.Id = int(len(r.operations) + 1)
		r.operations[int(oper.Id)] = oper
	}
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
