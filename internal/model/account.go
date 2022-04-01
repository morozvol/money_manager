package model

type Account struct {
	Id         int64    `db:"id""`
	Balance    float32  `db:"balance"`
	Currency   Currency `db:"currency"`
	Name       string   `db:"name"`
	Operations []Operation
	IdUser     int `db:"id_user"`
}
