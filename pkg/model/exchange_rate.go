package model

import "time"

type ExchangeRate struct {
	Id             int       `db:"id"               json:"id"`
	IdCurrencyFrom int       `db:"id_currency_from" json:"id_currency_from" validate:"required"`
	IdCurrencyTo   int       `db:"id_currency_to"   json:"id_currency_to"   validate:"required"`
	Rate           float32   `db:"rate"             json:"rate"             validate:"required"`
	Date           time.Time `db:"date"             json:"date"             validate:"required"`
}
