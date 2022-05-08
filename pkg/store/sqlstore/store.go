package sqlstore

import (
	"github.com/jmoiron/sqlx"
	"github.com/morozvol/money_manager/pkg/store"
)

// Store ...
type Store struct {
	db                      *sqlx.DB
	userRepository          *UserRepository
	accountRepository       *AccountRepository
	operationRepository     *OperationRepository
	currencyRepository      *CurrencyRepository
	categoryRepository      *CategoryRepository
	categoryLimitRepository *CategoryLimitRepository
	exchangeRateRepository  *ExchangeRateRepository
}

// New ...
func New(dbPool *sqlx.DB) store.Store {
	return &Store{
		db: dbPool,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Account ...
func (s *Store) Account() store.AccountRepository {
	if s.accountRepository != nil {
		return s.accountRepository
	}

	s.accountRepository = &AccountRepository{
		store: s,
	}

	return s.accountRepository
}

// Operation ...
func (s *Store) Operation() store.OperationRepository {
	if s.operationRepository != nil {
		return s.operationRepository
	}

	s.operationRepository = &OperationRepository{
		store: s,
	}

	return s.operationRepository
}

// Currency ...
func (s *Store) Currency() store.CurrencyRepository {
	if s.currencyRepository != nil {
		return s.currencyRepository
	}

	s.currencyRepository = &CurrencyRepository{
		store: s,
	}

	return s.currencyRepository
}

func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store: s,
	}

	return s.categoryRepository
}

func (s *Store) CategoryLimit() store.CategoryLimitRepository {
	if s.categoryLimitRepository != nil {
		return s.categoryLimitRepository
	}

	s.categoryLimitRepository = &CategoryLimitRepository{
		store: s,
	}

	return s.categoryLimitRepository
}

func (s *Store) ExchangeRate() store.ExchangeRateRepository {
	if s.exchangeRateRepository != nil {
		return s.exchangeRateRepository
	}

	s.exchangeRateRepository = &ExchangeRateRepository{
		store: s,
	}

	return s.exchangeRateRepository
}
