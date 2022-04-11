package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestCurrencyRepository_Find(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate()
	tests := []struct {
		name    string
		store   *Store
		id      int
		want    *model.Currency
		wantErr bool
	}{
		{
			"find USD",
			store,
			1,
			&model.Currency{Id: 1, Code: "USD", Name: "Долар США"},
			false,
		},
		{
			"find non exist",
			store,
			100,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CurrencyRepository{
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

func TestCurrencyRepository_GetAll(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate()
	tests := []struct {
		name    string
		store   *Store
		wantErr bool
	}{
		{
			"get all",
			store,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CurrencyRepository{
				store: tt.store,
			}
			_, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
