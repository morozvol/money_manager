package model

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Accounts []Account
}
