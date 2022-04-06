package tgbot

import (
	"fmt"
	objs "github.com/SakoDroid/telego/objects"
	exr "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"github.com/morozvol/money_manager/internal/model"
)

func (bot *tgbot) addOperation(u *objs.Update) {
	chatId := u.Message.Chat.Id
	userId := u.Message.From.Id

	_, err := bot.store.User().Find(userId)
	if err != nil {
		bot.help(u)
		return
	}
	uc := &userChat{userId, chatId}
	operation := model.Operation{}
	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.chatId)

	account := bot.accountsKeyboard(uc, msgChannel, msgEditor)
	operation.IdAccount = account.Id
	operation.Category = *bot.categoriesKeyboard(uc, msgChannel, msgEditor)
	currency := bot.choosePaymentCurrency(uc, msgChannel, msgEditor)
	operation.Sum = bot.getFloat(uc, msgChannel, "Введите сумму")

	ex := swap.NewSwap()

	ex.AddExchanger(exr.NewYahooApi(nil)).Build()

	rate := ex.Latest(fmt.Sprintf("%s/%s", currency.Code, account.Currency.Code)).GetRateValue()

	rate = rate + (rate*3)/100

	operation.Sum = operation.Sum * float32(rate)
	err = bot.store.Operation().Create(&operation)
	if err != nil {
		bot.Logger.Error(err.Error())
		bot.sendText(chatId, "Ошибка. На счету недостаточно средств")
	}
}
