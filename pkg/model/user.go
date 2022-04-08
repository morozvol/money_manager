package model

type User struct {
	Id                int64  `db:"id"`
	Name              string `db:"name"`
	DefaultCurrencyId int    `db:"id_default_currency"`
	Accounts          []Account
}
