package model

type AccountTypes int64

const (
	Card AccountTypes = iota + 1
	Cash
	Saving
)

func (at AccountTypes) ToString() string {
	return []string{"Card", "Cash", "Saving"}[at]
}

type AccountType struct {
	Id          AccountTypes `db:"id"`
	Name        string       `db:"name"`
	Symbol      string       `db:"symbol"`
	Description string       `db:"description"`
}
