package tgbot

import (
	objs "github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"time"
)

func (bot *tgbot) register(uc *objects.UserChat) {
	_, err := bot.store.User().Find(uc.UserId)
	if err == store.ErrRecordNotFound {
		user := &model.User{Id: uc.UserId, Name: ""}
		bot.store.User().Create(user)
	}
}

func (bot *tgbot) help(u *objs.Update) {
	uc := &objects.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}
	if !bot.beforeExecution(uc) {
		return
	}
	chatId := u.Message.Chat.Id
	helpMessage := "/add_account - добавить кошелёк\n" +
		"/add_operation - Новая операция\n" +
		"/info - информация о доступном баллансе"
	bot.sendTemporaryText(chatId, helpMessage, time.Second*10)
}
