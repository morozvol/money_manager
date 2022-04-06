package tgbot

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
)

func (bot *tgbot) currencyKeyboard(uc *userChat, messageChannel chan string, editor *bt.MessageEditor) *model.Currency {
	currencies, err := bot.store.Currency().GetAll()
	if err != nil {
		bot.error(err, "currencyKeyboard: не удалось получить валюты из db", nil)
	}

	kb := bot.CreateInlineKeyboard()
	for i, currency := range currencies {
		kb.AddCallbackButton(currency.Code,
			fmt.Sprintf("id currency: %d", currency.Id), i+1)
	}

	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, "Выбор валюты", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.error(err, "currencyKeyboard: не удалось отправить сообщение", nil)
	}
	defer func() {
		_, err := editor.DeleteMessage(msg.Result.MessageId)
		if err != nil {
			bot.error(err, "currencyKeyboard: не удалось удалить сообщение", msg)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id currency", messageChannel, ctx)

	if val, err := strconv.ParseInt(<-messageChannel, 10, 64); err == nil {
		for _, c := range currencies {
			if c.Id == val {
				return &c
			}
		}
	}
	return nil
}
