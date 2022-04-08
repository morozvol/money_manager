package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		u *model.User
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
			r := &UserRepository{
				store: tt.fields.store,
			}
			if err := r.Create(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_Find(t *testing.T) {
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
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
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

func Test_user_toModel(t *testing.T) {
	type fields struct {
		Id                int64
		Name              string
		DefaultCurrencyId sql.NullInt64
	}
	tests := []struct {
		name   string
		fields fields
		want   model.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := user{
				Id:                tt.fields.Id,
				Name:              tt.fields.Name,
				DefaultCurrencyId: tt.fields.DefaultCurrencyId,
			}
			if got := u.toModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
