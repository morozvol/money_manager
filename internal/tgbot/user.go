package tgbot

import (
	objs "github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

func (bot *tgbot) register(u *objs.Update) {
	userData := u.Message.From
	chatId := u.Message.Chat.Id
	_, err := bot.store.User().Find(userData.Id)
	if err == store.ErrRecordNotFound {
		user := &model.User{Id: int64(userData.Id), Name: userData.Username}
		err = bot.store.User().Create(user)
		if err != nil {
			bot.error(err, "пользователь не может быть зарегистрирован: ", userData)
			bot.sendText(chatId, "вы уже зарегистрированы!!!")
		}
		bot.sendText(chatId, "вы успешно зарегистрированы")

	} else {
		bot.sendText(chatId, "вы уже зарегистрированы")
	}
}

func (bot *tgbot) help(u *objs.Update) {
	helpMessage := ""
	chatId := u.Message.Chat.Id
	_, err := bot.store.User().Find(u.Message.From.Id)
	if err == store.ErrRecordNotFound {
		helpMessage = helpMessage + "первое что требуется сделать для настройки бота - регистрация. Пройти её можно тут /register.\n" +
			"Все остальные действия делаются только после регистрации!!!\n"
	}
	helpMessage = helpMessage + "/add_account - добавить кошелёк\n" +
		"/add_operation - Новая операция\n" +
		"/info - информация о доступном баллансе"

	bot.sendText(chatId, helpMessage)
}
