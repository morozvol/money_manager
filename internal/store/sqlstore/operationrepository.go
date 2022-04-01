package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

type OperationRepository struct {
	store *Store
}

// Create ...
func (r *OperationRepository) Create(o *model.Operation) error {

	_, err := r.store.db.Exec(
		"CALL public.apply_operation($1, $2, $3)",
		o.IdAccount,
		o.Sum,
		o.OperationType,
	)
	return err
}

// Find ...
func (r *OperationRepository) Find(id int) (*model.Operation, error) {
	u := &model.Operation{}
	if err := r.store.db.QueryRowx(
		"SELECT id, time, sum FROM operation WHERE id = $1",
		id,
	).StructScan(u); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
