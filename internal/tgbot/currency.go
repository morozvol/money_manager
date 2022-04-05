package tgbot

import (
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
)

func (bot *tgbot) currencyKeyboard(uc *userChat) *model.Currency {
	currencies, err := bot.store.Currency().GetAll()
	if err != nil {
		bot.Error(err, "currencyKeyboard: не удалось получить валюты из db", nil)
	}

	kb := bot.CreateInlineKeyboard()
	for i, currency := range currencies {
		kb.AddCallbackButton(currency.Code,
			fmt.Sprintf("id currency: %d", currency.Id), i+1)
	}

	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, "Выбор валюты", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.Error(err, "currencyKeyboard: не удалось отправить сообщение", nil)
	}
	defer func(editor *bt.MessageEditor, messageId int) {
		_, err := editor.DeleteMessage(messageId)
		if err != nil {
			bot.Error(err, "currencyKeyboard: не удалось удалить сообщение", msg)

		}
	}(bot.GetMsgEditor(uc.chatId), msg.Result.MessageId)

	if val, err := strconv.ParseInt(<-bot.userData[*uc], 10, 64); err == nil {
		for _, c := range currencies {
			if c.Id == val {
				return &c
			}
		}
	}
	return nil
}
