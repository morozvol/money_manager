package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

type CurrencyRepository struct {
	store *Store
}

// Find ...
func (r *CurrencyRepository) Find(id int) (*model.Currency, error) {
	c := &model.Currency{}
	if err := r.store.db.QueryRowx(
		"SELECT id,code,name FROM currency WHERE id = $1",
		id,
	).StructScan(c); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return c, nil
}
func (r *CurrencyRepository) GetAll() ([]model.Currency, error) {
	c := model.Currency{}
	res := []model.Currency{}
	rows, err := r.store.db.Queryx("SELECT id,code,name FROM currency")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}
