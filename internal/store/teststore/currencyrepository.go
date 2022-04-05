package teststore

import (
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

type CurrencyRepository struct {
	store      *Store
	currencies map[int]*model.Currency
}

// Find ...
func (r *CurrencyRepository) Find(id int) (*model.Currency, error) {
	o, ok := r.currencies[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return o, nil
}

func (r *CurrencyRepository) GetAll() ([]model.Currency, error) {
	var res []model.Currency

	for _, c := range r.currencies {
		res = append(res, *c)
	}
	return res, nil
}
