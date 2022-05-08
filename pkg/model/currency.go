package model

type Currency struct {
	Id   int    `db:"id"   json:"id"`
	Code string `db:"code" json:"code" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}
