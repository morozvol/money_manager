package exchange

import (
	"fmt"
	exr "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"github.com/morozvol/money_manager/pkg/model"
)

var MainCurrency = &model.Currency{Id: 1, Code: "USD"}

func exchange(currencyFrom *model.Currency, currencyTo *model.Currency) (rate float32, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = fmt.Errorf("panic occurred exchange: %v", er)
		}
	}()

	if currencyFrom.Id != currencyTo.Id {
		ex := swap.NewSwap()
		ex.AddExchanger(exr.NewYahooApi(nil)).
			AddExchanger(exr.NewTheMoneyConverterApi(nil)).
			Build()
		rate = float32(ex.Latest(fmt.Sprintf("%s/%s", currencyFrom.Code, currencyTo.Code)).GetRateValue())

		return
	}
	rate = 1
	return
}

func Exchange(currencyFrom *model.Currency, account *model.Account) (rate float32, err error) {

	rate, err = exchange(currencyFrom, &account.Currency)
	if err != nil {

		rate1, err1 := exchange(currencyFrom, MainCurrency)
		if err1 != nil {
			err = err1
			return
		}
		rate2, err2 := exchange(MainCurrency, &account.Currency)
		if err2 != nil {
			err = err2
			return
		}
		err = nil
		rate = rate1 * rate2
	}

	if account.AccountType.Id.ToString() == "Card" {
		rate = rate + (rate*3)/100
	}
	return
}
