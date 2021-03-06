package tgbot

import (
	"context"
	"fmt"
	"github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"strconv"
	"strings"
	"time"
)

func (bot *tgbot) sendText(chatId int, text string) {
	_, err := bot.SendMessage(chatId, text, "", 0, false, false)
	if err != nil {
		bot.error(err, "sendText: "+ErrSendMessage.Error(), nil)
	}
}

func (bot *tgbot) sendTemporaryText(chatId int, text string, dur time.Duration) {
	msg, err := bot.SendMessage(chatId, text, "", 0, false, false)
	if err != nil {
		bot.error(err, "sendText: "+ErrSendMessage.Error(), nil)
		return
	}
	time.Sleep(dur)
	editor := bot.GetMsgEditor(chatId)
	_, err = editor.DeleteMessage(msg.Result.MessageId)
	if err != nil {
		return
	}
}

func (bot *tgbot) error(err error, msg string, data interface{}) {
	if data == nil {
		bot.Logger.Error(fmt.Errorf("%w\n\t%s.", err, msg).Error())
	} else {
		bot.Logger.Error(fmt.Errorf("%w\n\t%s. (%#v)", err, msg, data).Error())
	}
	bot.Logger.Error(fmt.Errorf(msg+" %w %#v", err, data).Error())
}

func getCurrencyById(id int, currencies []model.Currency) model.Currency {
	for _, item := range currencies {
		if item.Id == id {
			return item
		}
	}
	return model.Currency{}
}
func (bot *tgbot) getFloat(uc *o.UserChat, messageChannel chan string, text string, parentCtx context.Context) (float32, error) {
	bot.sendText(uc.ChatId, text)

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.registerChannel(uc, "message", "", messageChannel, ctx)

	return getFloatFromChannel(messageChannel, ctx)
}
func (bot *tgbot) getString(uc *o.UserChat, messageChannel chan string, text string, parentCtx context.Context) (string, error) {
	bot.sendText(uc.ChatId, text)

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.registerChannel(uc, "message", "", messageChannel, ctx)
	val, err := getStringFromChannel(messageChannel, ctx)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (bot *tgbot) registerChannel(uc *o.UserChat, mediaType string, callbackInfo string, data chan<- string, ctx context.Context) {
	messageChannel, err := bot.AdvancedMode().RegisterChannel(strconv.Itoa(uc.ChatId), mediaType)
	if err != nil {
		bot.Logger.Fatal(err.Error())
	}
	var ar []string

	for {
		select {
		case up := <-*messageChannel:
			{
				switch up.GetType() {
				case "message":
					if up.Message.From.Id == uc.UserId {
						data <- up.Message.Text
					}
				case "callback_query":
					ar = strings.Split(up.CallbackQuery.Data, ": ")
					if up.CallbackQuery.From.Id == uc.UserId && ar[0] == callbackInfo {
						data <- ar[1]
					}
				}
			}
		case <-ctx.Done():
			{
				bot.AdvancedMode().UnRegisterChannel(strconv.Itoa(uc.ChatId), mediaType)
				return
			}
		}
	}
}

func (bot *tgbot) cancelOperation(u *objects.Update) {
	uc := o.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}
	bot.taskCancel.Cancel(uc)
}
func (bot *tgbot) tryCancelOperation(uc *o.UserChat) bool {
	_, ok := bot.taskCancel.Load(*uc)
	if ok {
		bot.sendTemporaryText(uc.ChatId, "?? ???????????? ???????????? ???????????????????? ???????????? ????????????????. ?????? ???? ???????????? ?????????????? /cancel", time.Second*5)

		_, ok := bot.taskCancel.Load(*uc)
		if ok {
			return false
		}
		return true
	}
	return true
}

func (bot tgbot) sendInlineKeyboard(uc *o.UserChat, text string, kb telego.MarkUps) (*objects.Message, error) {
	msg, err := bot.AdvancedMode().ASendMessage(uc.ChatId, text, "", 0, false, false, nil, false, false, kb)
	if err != nil {
		return nil, err
	}
	return msg.Result, nil

}

func (bot tgbot) beforeExecution(uc *o.UserChat) bool {
	_, err := bot.store.User().Find(uc.UserId)
	if err == store.ErrRecordNotFound {
		bot.register(uc)
	}
	return bot.tryCancelOperation(uc)

}

func getIntFromChannel(messageChannel <-chan string, ctx context.Context) (int, error) {
	select {
	case msg := <-messageChannel:
		if val, err := strconv.ParseInt(msg, 10, 64); err == nil {
			return int(val), nil
		}
		return 0, ErrUnknown
	case <-ctx.Done():
		return 0, ErrOperationCanceled
	}
}

func getStringFromChannel(messageChannel <-chan string, ctx context.Context) (string, error) {
	select {
	case msg := <-messageChannel:
		return msg, nil
	case <-ctx.Done():
		return "", ErrOperationCanceled
	}
}

func getFloatFromChannel(messageChannel <-chan string, ctx context.Context) (float32, error) {
	select {
	case msg := <-messageChannel:
		msg = strings.Replace(msg, ",", ".", -1)
		val, err := strconv.ParseFloat(msg, 32)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	case <-ctx.Done():
		return 0, ErrOperationCanceled
	}
}

func getTwoIntFromChannel(messageChannel <-chan string, ctx context.Context) (int, int, error) {
	select {
	case msg := <-messageChannel:
		ar := strings.Split(msg, " ")
		if val1, err := strconv.ParseInt(ar[0], 10, 64); err == nil {
			if val2, err := strconv.ParseInt(ar[1], 10, 64); err == nil {
				return int(val1), int(val2), nil
			}
		}
		return 0, 0, ErrUnknown
	case <-ctx.Done():
		return 0, 0, ErrOperationCanceled
	}
}
