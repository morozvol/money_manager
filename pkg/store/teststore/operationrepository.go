package teststore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"time"
)

type OperationRepository struct {
	store      *Store
	operations map[int]*model.Operation
}

// Create ...
func (r *OperationRepository) Create(o ...*model.Operation) error {
	for _, oper := range o {
		oper.Id = len(r.operations) + 1
		r.operations[oper.Id] = oper
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

func (r *OperationRepository) Get(timeFrom, timeTo time.Time) ([]model.Operation, error) {

	return nil, nil
}
