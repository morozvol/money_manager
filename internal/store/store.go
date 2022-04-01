package store

// Store ...
type Store interface {
	User() UserRepository
	Account() AccountRepository
	Operation() OperationRepository
	Currency() CurrencyRepository
}
