package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
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
			model.GetAccount100(int(u.Id), model.Card, model.Currency{Id: 1}),
			false,
		},
		{
			"valid duplicate",
			store,
			model.GetAccount100(int(u.Id), model.Card, model.Currency{Id: 1}),
			false,
		},
		{
			"invalid currency non exist",
			store,
			model.GetAccount100(int(u.Id), model.Card, model.Currency{Id: 15}),
			true,
		},
		{
			"invalid type non exist",
			store,
			model.GetAccount100(int(u.Id), 15, model.Currency{Id: 1}),
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

	tests := []struct {
		name    string
		store   *Store
		id      int
		want    *model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.store,
			}
			got, err := r.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_FindByUserId(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.fields.store,
			}
			got, err := r.FindByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
