package tgbot

import (
	"context"
	objs "github.com/SakoDroid/telego/objects"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/core/exchange"
	"github.com/morozvol/money_manager/pkg/core/system_category"
	model "github.com/morozvol/money_manager/pkg/model"
)

func (bot *tgbot) addTransferOperation(u *objs.Update) {

	uc := &o.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}

	ctx, cancel := context.WithCancel(context.Background())

	bot.taskCancel.Store(*uc, cancel)

	_, err := bot.store.User().Find(uc.UserId)
	if err != nil {
		bot.help(u)
		return
	}

	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.ChatId)

	accountFrom, err := bot.accountsKeyboard(uc, msgChannel, msgEditor, ctx)
	if err != nil {
		return
	}

	accountTo, err := bot.accountsKeyboard(uc, msgChannel, msgEditor, ctx)
	if err != nil {
		return
	}

	sum, err := bot.getFloat(uc, msgChannel, "Введите сумму которую хотите перевести", ctx)
	if err != nil {
		return
	}
	sum = sum * exchange.Exchange(&accountTo.Currency, accountFrom)

	sc := system_category.GetCategory(bot.store)

	operationComing := model.Operation{Sum: sum, IdAccount: accountFrom.Id, Category: model.Category{Id: sc.IdConsumptionTransfer}}
	operationConsumption := model.Operation{Sum: sum, IdAccount: accountTo.Id, Category: model.Category{Id: sc.IdComingTransfer}}

	err = bot.store.Operation().Create(&operationComing, &operationConsumption)
	if err != nil {
		bot.error(err, "addTransferOperation", nil)
		bot.sendText(uc.ChatId, "Ошибка. Не удалось выполнить операцию")
	}
}
