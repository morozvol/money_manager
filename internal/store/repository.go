package store

import (
	"github.com/morozvol/money_manager/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

// AccountRepository ...
type AccountRepository interface {
	Create(account *model.Account) error
	Find(int) (*model.Account, error)
	FindByUserId(int) ([]model.Account, error)
}

// OperationRepository ...
type OperationRepository interface {
	Create(operation *model.Operation) error
	Find(int) (*model.Operation, error)
}

// CurrencyRepository ...
type CurrencyRepository interface {
	Find(int) (*model.Currency, error)
	GetAll() ([]model.Currency, error)
}
