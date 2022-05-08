package model

type User struct {
	Id                int       `db:"id"                  json:"id"                  validate:"required"`
	Name              string    `db:"name"                json:"name"                validate:"required"`
	DefaultCurrencyId int       `db:"id_default_currency" json:"id_default_currency"`
	Accounts          []Account `                         json:"accounts,omitempty"`
}
