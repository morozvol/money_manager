package exchange

import (
	"fmt"
	exr "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"time"
)

var MainCurrency = &model.Currency{Id: 1, Code: "USD"}

func getRate(currencyFrom *model.Currency, currencyTo *model.Currency, s store.Store) (rate float32, err error) {
	rate, err = s.ExchangeRate().Get(currencyFrom.Id, currencyTo.Id, time.Now())
	if err != nil {

		rate, err = exchange(currencyFrom, currencyTo)
		if err != nil {

			rate1, err1 := exchange(currencyFrom, MainCurrency)
			if err1 != nil {
				err = err1
				return
			}
			rate2, err2 := exchange(MainCurrency, currencyTo)
			if err2 != nil {
				err = err2
				return
			}
			err = nil
			rate = rate1 * rate2
		}
		err := s.ExchangeRate().Create(&model.ExchangeRate{IdCurrencyFrom: currencyFrom.Id, IdCurrencyTo: currencyTo.Id, Rate: rate, Date: time.Now()})
		if err != nil {
			return 0, err
		}
	}
	return
}

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

func Exchange(currencyFrom *model.Currency, account *model.Account, s store.Store) (rate float32, err error) {
	rate, err = getRate(currencyFrom, &account.Currency, s)
	if err != nil {
		return 0, err
	}

	if account.AccountType.Id.ToString() == "Card" {
		rate = rate + (rate*3)/100
	}
	return
}
