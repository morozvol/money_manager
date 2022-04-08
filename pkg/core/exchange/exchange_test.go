package exchange

import (
	"github.com/morozvol/money_manager/pkg/model"
	"testing"
)

func TestExchange(t *testing.T) {
	type args struct {
		currencyFrom *model.Currency
		account      *model.Account
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exchange(tt.args.currencyFrom, tt.args.account); got == 0 {
				t.Errorf("Exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
