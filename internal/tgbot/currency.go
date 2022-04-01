package tgbot

import (
	"fmt"
	objs "github.com/morozvol/telego/objects"
)

func (bot *tgbot) currencyKeyboard(u *objs.Update) {
	currencies, err := bot.store.Currency().GetAll()
	if err != nil {
		bot.Logger.Error(err.Error())
	}

	kb := bot.CreateInlineKeyboard()
	for i, currency := range currencies {
		kb.AddCallbackButton(currency.Code,
			fmt.Sprintf("id currency: %d", currency.Id), i+1)
	}

	_, err = bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "Выбор валюты", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		fmt.Println(err)
	}
}
