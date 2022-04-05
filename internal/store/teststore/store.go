package teststore

import (
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

// Store ...
type Store struct {
	userRepository      *UserRepository
	operationRepository *OperationRepository
	currencyRepository  *CurrencyRepository
	categoryRepository  *CategoryRepository
	accountRepository   *AccountRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
func (s *Store) Account() store.AccountRepository {
	if s.accountRepository != nil {
		return s.accountRepository
	}

	s.accountRepository = &AccountRepository{
		store:    s,
		accounts: make(map[int]*model.Account),
	}

	return s.accountRepository
}

// Operation ...
func (s *Store) Operation() store.OperationRepository {
	if s.operationRepository != nil {
		return s.operationRepository
	}

	s.operationRepository = &OperationRepository{
		store:      s,
		operations: make(map[int]*model.Operation),
	}
	return s.operationRepository
}

// Currency ...
func (s *Store) Currency() store.CurrencyRepository {
	if s.currencyRepository != nil {
		return s.currencyRepository
	}

	s.currencyRepository = &CurrencyRepository{
		store:      s,
		currencies: make(map[int]*model.Currency),
	}

	return s.currencyRepository
}

func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store:      s,
		categories: make(map[int]*model.Category),
	}

	return s.categoryRepository
}
