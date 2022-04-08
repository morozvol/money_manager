package sqlstore

import (
	"context"
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

type OperationRepository struct {
	store *Store
}

// Create ...
func (r *OperationRepository) Create(o ...*model.Operation) error {

	ctx := context.Background()
	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, operation := range o {
		_, err = tx.ExecContext(ctx, "CALL public.apply_operation($1, $2, $3, $4)",
			operation.IdAccount,
			operation.Sum,
			operation.Category.Id,
			operation.Description)
		if err != nil {
			err_rb := tx.Rollback()
			if err_rb != nil {
				return err_rb
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
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
