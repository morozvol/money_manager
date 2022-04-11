package tgbot

import (
	"context"
	"fmt"
	objs "github.com/SakoDroid/telego/objects"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/core/exchange"
	"github.com/morozvol/money_manager/pkg/model"
)

func (bot *tgbot) addPaymentOperation(u *objs.Update) {
	bot.beforeExecution(u)

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
	rate, err := exchange.Exchange(currency, account)
	if err != nil {
		bot.sendText(uc.ChatId, "Ошибка. Не удалось получить курс "+currency.Code+"/"+account.Currency.Code+".")
		bot.error(err, "addPaymentOperation", nil)
		return
	}
	bot.Logger.Debug(fmt.Sprintf("курс: %f", rate))
	operation.Sum = sum * rate

	err = bot.store.Operation().Create(&operation)
	if err != nil {
		bot.Logger.Error(err.Error())
		bot.sendText(uc.ChatId, "Ошибка. На счету недостаточно средств")
		return
	}
	bot.successOperation(uc, &operation)
}

func (bot *tgbot) successOperation(uc *o.UserChat, operation *model.Operation) {
	tf := bot.GetTextFormatter()
	tf.AddBold("Операция успешно выполнена.\n")
	if operation.Category.Type == model.Coming {
		tf.AddNormal(fmt.Sprintf("Пополнено %.2f", operation.Sum))
	} else {
		tf.AddNormal(fmt.Sprintf("Cписано %.2f", operation.Sum))
	}

	bot.sendText(uc.ChatId, tf.GetText())
}
