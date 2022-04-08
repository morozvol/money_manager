package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestAccountRepository_Create(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		a *model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.fields.store,
			}
			if err := r.Create(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepository_Find(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				store: tt.fields.store,
			}
			got, err := r.Find(tt.args.id)
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
