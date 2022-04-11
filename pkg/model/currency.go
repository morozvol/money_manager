package model

type Currency struct {
	Id   int64  `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}
