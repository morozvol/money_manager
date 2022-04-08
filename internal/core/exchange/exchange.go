package exchange

import (
	"fmt"
	exr "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"github.com/morozvol/money_manager/internal/model"
)

func Exchange(currencyFrom *model.Currency, account *model.Account) float32 {
	if currencyFrom.Id != account.Currency.Id {
		ex := swap.NewSwap()
		ex.AddExchanger(exr.NewYahooApi(nil)).Build()
		rate := ex.Latest(fmt.Sprintf("%s/%s", currencyFrom.Code, account.Currency.Code)).GetRateValue()

		if account.AccountType.Id.ToString() == "Card" {
			rate = rate + (rate*3)/100
		}
		return float32(rate)
	}
	return 1
}
