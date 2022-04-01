package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/internal/model"
	"github.com/morozvol/money_manager/internal/store"
)

type AccountRepository struct {
	store *Store
}

// Create ...
func (r *AccountRepository) Create(a *model.Account) error {
	return r.store.db.QueryRowx(
		"INSERT INTO account (name, balance, id_currency, id_user) VALUES ($1,$2,$3,$4)",
		a.Name,
		a.Balance,
		a.Currency.Id,
		a.IdUser,
	).Err()
}

// Find ...
func (r *AccountRepository) Find(id int) (*model.Account, error) {
	a := &model.Account{}
	if err := r.store.db.QueryRowx(
		"SELECT id as id_account, balance, id_currency, name as account_name FROM account WHERE id = $1",
		id,
	).StructScan(a); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return a, nil
}

// FindByUserId ...
func (r *AccountRepository) FindByUserId(userId int) ([]model.Account, error) {

	q := "SELECT id, balance, id_currency as \"currency.id\", name FROM account WHERE id_user = $1"

	rows, err := r.store.db.Queryx(q, userId)
	if err != nil {
		return nil, err
	}

	var accounts []model.Account
	var account model.Account
	for rows.Next() {
		err := rows.StructScan(&account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
