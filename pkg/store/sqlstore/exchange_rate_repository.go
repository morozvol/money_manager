package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"time"
)

type ExchangeRateRepository struct {
	store *Store
}

func (r ExchangeRateRepository) Get(idCurrencyFrom, idCurrencyTo int, time time.Time) (float32, error) {
	a := &model.ExchangeRate{}
	if err := r.store.db.QueryRowx(
		"SELECT id, id_currency_from, id_currency_to, rate, date FROM exchange_rate WHERE id_currency_from = $1 AND id_currency_to = $2 AND date = $3::date",
		idCurrencyFrom, idCurrencyTo, time,
	).StructScan(a); err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}
	return a.Rate, nil
}

func (r ExchangeRateRepository) Create(rate *model.ExchangeRate) error {
	lastInsertId := 0
	err := r.store.db.QueryRow(
		"INSERT INTO exchange_rate (id_currency_from, id_currency_to, rate, date) VALUES ($1,$2,$3,$4) RETURNING id;",
		rate.IdCurrencyFrom,
		rate.IdCurrencyTo,
		rate.Rate,
		rate.Date,
	).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	rate.Id = lastInsertId
	return nil
}
