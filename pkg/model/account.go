package model

type Account struct {
	Id          int      `db:"id"`
	Balance     float32  `db:"balance"`
	Currency    Currency `db:"currency"`
	Name        string   `db:"name"`
	Operations  []Operation
	IdUser      int         `db:"id_user"`
	AccountType AccountType `db:"account_type"`
}
