package model

type CategoryLimit struct {
	Id         int     `db:"id"          json:"id"`
	IdUser     int     `db:"id_user"     json:"id_user"     validate:"required"`
	IdCategory int     `db:"id_category" json:"id_category" validate:"required"`
	IdCurrency int     `db:"id_currency" json:"id_currency" validate:"required"`
	Sum        float32 `db:"sum"         json:"sum"         validate:"gt=0"`
}
