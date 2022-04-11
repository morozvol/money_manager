package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

type AccountRepository struct {
	store *Store
}

// Create ...
func (r *AccountRepository) Create(a *model.Account) error {
	lastInsertId := 0
	err := r.store.db.QueryRow(
		"INSERT INTO account (name, balance, id_currency, id_user, id_account_type) VALUES ($1,$2,$3,$4,$5) RETURNING id;",
		a.Name,
		a.Balance,
		a.Currency.Id,
		a.IdUser,
		a.AccountType.Id,
	).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	a.Id = int(lastInsertId)
	return nil
}

// Find ...
func (r *AccountRepository) Find(id int) (*model.Account, error) {
	a := &model.Account{}
	if err := r.store.db.QueryRowx(
		"SELECT a.id, a.balance, a.id_currency as \"currency.id\", a.name, a.id_account_type as \"account_type.id\", t.symbol as \"account_type.symbol\" FROM account a LEFT JOIN account_type t on a.id_account_type = t.id WHERE a.id = $1",
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

	q := "SELECT a.id, a.balance, a.id_currency as \"currency.id\", a.name, a.id_account_type as \"account_type.id\", t.symbol as \"account_type.symbol\" FROM account a LEFT JOIN account_type t on a.id_account_type = t.id WHERE id_user = $1"

	rows, err := r.store.db.Queryx(q, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
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
	if len(accounts) == 0 {
		return nil, store.ErrRecordNotFound
	}

	return accounts, nil
}
