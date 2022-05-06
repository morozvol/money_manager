package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"testing"
)

func TestAccountRepository_Create(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("account", "\"user\"")
	u := model.GetUser()

	r := &UserRepository{
		store: store,
	}
	err := r.Create(u)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	tests := []struct {
		name    string
		store   *Store
		a       *model.Account
		wantErr bool
	}{
		{
			"valid",
			store,
			model.GetAccount100(u.Id, model.Card, model.Currency{Id: 1}),
			false,
		},
		{
			"valid duplicate",
			store,
			model.GetAccount100(u.Id, model.Card, model.Currency{Id: 1}),
			false,
		},
		{
			"invalid currency non exist",
			store,
			model.GetAccount100(u.Id, model.Card, model.Currency{Id: 15}),
			true,
		},
		{
			"invalid type non exist",
			store,
			model.GetAccount100(u.Id, 15, model.Currency{Id: 1}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.store,
			}
			if err := r.Create(tt.a); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepository_Find(t *testing.T) {

	store, truncate := GetTestDBStore(t)
	defer truncate("account", "category", "\"user\"")

	r := &UserRepository{
		store: store,
	}
	ar := &AccountRepository{
		store: store,
	}

	u := model.GetUser()

	ac := model.GetAccount100(u.Id, model.Card, model.Currency{Id: 1})

	err := r.Create(u)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	err = ar.Create(ac)
	if err != nil {
		t.Fatal("счёт не может быть создан")
	}

	tests := []struct {
		name    string
		store   *Store
		id      int
		wantErr bool
	}{
		{
			"find valid account",
			store,
			1,
			false,
		},
		{
			"find invalid account",
			store,
			100,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.store,
			}
			_, err := r.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAccountRepository_FindByUserId(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("account", "\"user\"")
	u := model.GetUser()
	u1 := model.GetUser()
	u1.Id = 5
	ac := model.GetAccount100(u.Id, model.Card, model.Currency{Id: 1})

	r := &UserRepository{
		store: store,
	}
	err := r.Create(u)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}
	err = r.Create(u1)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	ar := &AccountRepository{
		store: store,
	}

	err = ar.Create(ac)
	if err != nil {
		t.Fatal("счёт не может быть создан")
	}

	tests := []struct {
		name    string
		store   *Store
		userId  int
		wantErr bool
	}{
		{
			"user with id non exist",
			store,
			10,
			true,
		},
		{
			"user without accounts",
			store,
			5,
			true,
		},
		{
			"user without accounts",
			store,
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.store,
			}
			_, err := r.FindByUserId(tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
