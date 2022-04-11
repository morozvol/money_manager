package exchange

import (
	"github.com/morozvol/money_manager/pkg/model"
	"testing"
)

func TestExchange(t *testing.T) {

	tests := []struct {
		name         string
		currencyFrom *model.Currency
		accountCard  *model.Account
		wantErr      bool
	}{
		{
			name:         "USD to EUR",
			currencyFrom: &model.Currency{Id: 1, Code: "USD", Name: ""},
			accountCard:  &model.Account{Id: 2, Currency: model.Currency{Id: 2, Code: "EUR"}, AccountType: model.AccountType{Id: model.Card}},
			wantErr:      false,
		},
		{
			name:         "BYN to TRY",
			currencyFrom: &model.Currency{Id: 4, Code: "BYN", Name: ""},
			accountCard:  &model.Account{Id: 2, Currency: model.Currency{Id: 2, Code: "TRY"}, AccountType: model.AccountType{Id: model.Card}},
			wantErr:      false,
		},
		{
			name:         "XXX to YYY",
			currencyFrom: &model.Currency{Id: 15, Code: "XXX", Name: ""},
			accountCard:  &model.Account{Id: 2, Currency: model.Currency{Id: 2, Code: "YYY"}, AccountType: model.AccountType{Id: model.Card}},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Exchange(tt.currencyFrom, tt.accountCard)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exchange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_exchange(t *testing.T) {

	tests := []struct {
		name         string
		currencyFrom *model.Currency
		currencyTo   *model.Currency
		wantErr      bool
	}{
		{
			name:         "USD to EUR",
			currencyFrom: &model.Currency{Id: 1, Code: "USD", Name: ""},
			currencyTo:   &model.Currency{Id: 2, Code: "EUR", Name: ""},
			wantErr:      false,
		},
		{
			name:         "BYN to TRY",
			currencyFrom: &model.Currency{Id: 4, Code: "BYN", Name: ""},
			currencyTo:   &model.Currency{Id: 3, Code: "TRY", Name: ""},
			wantErr:      true,
		},
		{
			name:         "XXX to YYY",
			currencyFrom: &model.Currency{Id: 15, Code: "XXX", Name: ""},
			currencyTo:   &model.Currency{Id: 16, Code: "YYY", Name: ""},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := exchange(tt.currencyFrom, tt.currencyTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("exchange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
