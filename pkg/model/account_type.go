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
	Id          AccountTypes `db:"id"          json:"id"          validate:"max=3,min=1"`
	Name        string       `db:"name"        json:"name"`
	Symbol      string       `db:"symbol"      json:"symbol"`
	Description string       `db:"description" json:"description"`
}
