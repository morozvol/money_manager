package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestCurrencyRepository_Find(t *testing.T) {
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
		want    *model.Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CurrencyRepository{
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

func TestCurrencyRepository_GetAll(t *testing.T) {
	type fields struct {
		store *Store
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CurrencyRepository{
				store: tt.fields.store,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
