package model

type User struct {
	Id                int    `db:"id"`
	Name              string `db:"name"`
	DefaultCurrencyId int    `db:"id_default_currency"`
	Accounts          []Account
}
