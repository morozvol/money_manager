package model

type Currency struct {
	Id   int    `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}
