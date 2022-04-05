package tgbot

import (
	"fmt"
	objs "github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
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
	bot.userData[*uc] = make(chan string)
	account := bot.accountsKeyboard(uc)

	operation.IdAccount = account.Id
	operation.Category = *bot.categoriesKeyboard(uc)

	bot.SendText(chatId, "Введите сумму")
	if fval, err := strconv.ParseFloat(<-bot.userData[*uc], 64); err == nil {
		operation.Sum = float32(fval)
	} else {
		bot.Error(err, "addOperation: Не удалось привести к float", nil)
	}

	close(bot.userData[*uc])
	delete(bot.userData, *uc)

	err = bot.store.Operation().Create(&operation)
	if err != nil {
		bot.Logger.Error(err.Error())
		bot.SendText(chatId, "Ошибка. На счету недостаточно средств")
	}
}

func (bot *tgbot) operationTypeKeyboard(u *objs.Update) int {

	kb := bot.CreateInlineKeyboard()
	kb.AddCallbackButton("Приход", fmt.Sprintf("id currency: %d", model.Coming), 1)
	kb.AddCallbackButton("Расход", fmt.Sprintf("id currency: %d", model.Consumption), 1)

	msg, err := bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "Выбор типа операции", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.Error(err, "addOperation: не удалось отправить сообшение", nil)
	}
	return msg.Result.MessageId
}
