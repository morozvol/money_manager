package model

type Account struct {
	Id          int         `db:"id" json:"id"`
	Balance     float32     `db:"balance" json:"balance"`
	Currency    Currency    `db:"currency" json:"currency"`
	Name        string      `db:"name" json:"name"`
	Operations  []Operation `json:"operations,omitempty"`
	IdUser      int         `db:"id_user" json:"id_user"`
	AccountType AccountType `db:"account_type" json:"account_type"`
}
