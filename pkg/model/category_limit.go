package model

type CategoryLimit struct {
	Id         int     `db:"id"`
	IdUser     int     `db:"id_user"`
	IdCategory int     `db:"id_category"`
	IdCurrency int     `db:"id_currency"`
	Sum        float32 `db:"sum"`
}
