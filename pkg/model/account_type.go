package model

type AccountTypes int

const (
	Card AccountTypes = iota + 1
	Cash
	Saving
)

func (at AccountTypes) ToString() string {
	return []string{"Card", "Cash", "Saving"}[at]
}

type AccountType struct {
	Id          AccountTypes `db:"id"          json:"id"`
	Name        string       `db:"name"        json:"name"        validate:"required"`
	Symbol      string       `db:"symbol"      json:"symbol"      validate:"required"`
	Description string       `db:"description" json:"description" validate:"required"`
}
