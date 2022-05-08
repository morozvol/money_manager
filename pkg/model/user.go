package model

type User struct {
	Id                int       `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	DefaultCurrencyId int       `db:"id_default_currency" json:"id_default_currency"`
	Accounts          []Account `json:"accounts,omitempty"`
}
