package tgbot

import (
	objs "github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"time"
)

func (bot *tgbot) register(u *objs.Update) {
	userData := u.Message.From
	chatId := u.Message.Chat.Id
	_, err := bot.store.User().Find(userData.Id)
	if err == store.ErrRecordNotFound {
		user := &model.User{Id: userData.Id, Name: userData.Username}
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
	bot.beforeExecution(u)

	chatId := u.Message.Chat.Id
	helpMessage := "/add_account - добавить кошелёк\n" +
		"/add_operation - Новая операция\n" +
		"/info - информация о доступном баллансе"
	bot.sendTemporaryText(chatId, helpMessage, time.Second*10)
}
