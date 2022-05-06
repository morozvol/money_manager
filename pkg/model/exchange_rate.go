package model

import "time"

type ExchangeRate struct {
	Id             int       `db:"id"`
	IdCurrencyFrom int       `db:"id_currency_from"`
	IdCurrencyTo   int       `db:"id_currency_to"`
	Rate           float32   `db:"rate"`
	Date           time.Time `db:"date"`
}
