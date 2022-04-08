package tgbot

import (
	"context"
	objs "github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/core/exchange"
	"github.com/morozvol/money_manager/internal/model"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
)

func (bot *tgbot) addOperation(u *objs.Update) {

	uc := &o.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}

	ctx, cancel := context.WithCancel(context.Background())

	bot.taskCancel.Store(*uc, cancel)

	user, err := bot.store.User().Find(uc.UserId)
	if err != nil {
		bot.help(u)
		return
	}
	operation := model.Operation{}
	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.ChatId)

	account, err := bot.accountsKeyboard(uc, msgChannel, msgEditor, ctx)
	if err != nil {
		return
	}
	operation.IdAccount = account.Id

	cat, err := bot.categoriesKeyboard(uc, msgChannel, msgEditor, ctx)
	if err != nil {
		return
	}
	operation.Category = *cat

	currency, err := bot.chooseCurrency(user, uc, msgChannel, msgEditor, ctx, true)
	if err != nil {
		return
	}

	sum, err := bot.getFloat(uc, msgChannel, "Введите сумму", ctx)
	if err != nil {
		return
	}
	operation.Sum = sum * exchange.Exchange(currency, account)

	err = bot.store.Operation().Create(&operation)
	if err != nil {
		bot.Logger.Error(err.Error())
		bot.sendText(uc.ChatId, "Ошибка. На счету недостаточно средств")
	}
}
