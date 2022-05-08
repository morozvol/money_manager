package model

import "time"

type ExchangeRate struct {
	Id             int       `db:"id"               json:"id"`
	IdCurrencyFrom int       `db:"id_currency_from" json:"id_currency_from"`
	IdCurrencyTo   int       `db:"id_currency_to"   json:"id_currency_to"`
	Rate           float32   `db:"rate"             json:"rate"`
	Date           time.Time `db:"date"             json:"date"`
}
