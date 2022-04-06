package tgbot

import (
	objs "github.com/SakoDroid/telego/objects"
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

	operation.IdAccount = bot.accountsKeyboard(uc, msgChannel, msgEditor).Id
	operation.Category = *bot.categoriesKeyboard(uc, msgChannel, msgEditor)
	operation.Sum = bot.getFloat(uc, msgChannel, "Введите сумму")

	err = bot.store.Operation().Create(&operation)
	if err != nil {
		bot.Logger.Error(err.Error())
		bot.sendText(chatId, "Ошибка. На счету недостаточно средств")
	}
}

/*func (bot *tgbot) operationTypeKeyboard(uc *userChat) int {

	kb := bot.CreateInlineKeyboard()
	kb.AddCallbackButton("Приход", fmt.Sprintf("id currency: %d", model.Coming), 1)
	kb.AddCallbackButton("Расход", fmt.Sprintf("id currency: %d", model.Consumption), 1)

	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, "Выбор типа операции", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.error(err, "addOperation: не удалось отправить сообшение", nil)
	}

	defer func(editor *bt.MessageEditor, messageId int) {
		_, err := editor.DeleteMessage(messageId)
		if err != nil {
			bot.error(err, "currencyKeyboard: не удалось удалить сообщение", msg)
		}
	}(bot.GetMsgEditor(uc.chatId), msg.Result.MessageId)

	messageChannel := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer close(messageChannel)
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id currency", messageChannel, ctx)
	return
}
*/
