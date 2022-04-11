package tgbot

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	model "github.com/morozvol/money_manager/pkg/model"
)

func (bot *tgbot) chooseCurrency(u *model.User, uc *o.UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context, isShowDefault bool) (*model.Currency, error) {
	shift := 1
	currencies, err := bot.store.Currency().GetAll()
	if err != nil {
		bot.error(err, "chooseCurrency: не удалось получить валюты из db", nil)
	}

	kb := bot.CreateInlineKeyboard()
	if u.DefaultCurrencyId != 0 && isShowDefault {
		kb.AddCallbackButton("По умолчанию",
			fmt.Sprintf("id currency: %d", u.DefaultCurrencyId), shift)
		shift++
	}

	for i, currency := range currencies {
		kb.AddCallbackButton(currency.Code,
			fmt.Sprintf("id currency: %d", currency.Id), int(i/2)+shift)
	}

	msg, err := bot.sendInlineKeyboard(uc, "Выбор валюты оплаты", kb)
	if err != nil {
		bot.error(err, "chooseCurrency: не удалось отправить сообщение", nil)
	}
	defer func() {
		_, err := editor.DeleteMessage(msg.MessageId)
		if err != nil {
			bot.error(err, "chooseCurrency: не удалось удалить сообщение", msg)
		}
	}()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id currency", messageChannel, ctx)

	val, err := getIntFromChannel(messageChannel, ctx)
	if err == nil {
		for _, c := range currencies {
			if c.Id == int64(val) {
				return &c, nil
			}
		}
	}
	return nil, err
}
