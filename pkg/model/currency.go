package model

type Currency struct {
	Id   int    `db:"id"   json:"id"`
	Code string `db:"code" json:"id_account"`
	Name string `db:"name"`
}
