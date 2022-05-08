package model

type CategoryLimit struct {
	Id         int     `db:"id" json:"id"`
	IdUser     int     `db:"id_user" json:"id_user"`
	IdCategory int     `db:"id_category" json:"id_category"`
	IdCurrency int     `db:"id_currency" json:"id_currency"`
	Sum        float32 `db:"sum" json:"sum"`
}
