package store

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/model/category_tree"
	"time"
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
	Create(operation ...*model.Operation) error
	Find(int) (*model.Operation, error)
	Get(time.Time, time.Time, int) ([]model.Operation, error)
}

// CurrencyRepository ...
type CurrencyRepository interface {
	Find(int) (*model.Currency, error)
	GetAll() ([]model.Currency, error)
}

// CategoryRepository ...
type CategoryRepository interface {
	Create(category *model.Category) error
	Get(int) (*category_tree.CategoryTree, error)
	GetSystem() ([]model.Category, error)
	Delete(int) error
}
type CategoryLimitRepository interface {
}
type ExchangeRateRepository interface {
	Get(int, int, time.Time) (float32, error)
	Create(rate *model.ExchangeRate) error
}
