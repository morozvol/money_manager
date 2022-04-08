package tgbot

import (
	"context"
	"errors"
	"fmt"
	"github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
	"strings"
)

var ErrOperationCanceled = errors.New("operation was cancel")
var ErrUnknown = errors.New("error")

func (bot *tgbot) sendText(chatId int, text string) {
	_, err := bot.SendMessage(chatId, text, "", 0, false, false)
	if err != nil {
		bot.error(err, "sendText: Не удалось отправить сообщение", nil)
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

func getCurrencyById(id int64, currencies []model.Currency) model.Currency {
	for _, item := range currencies {
		if item.Id == id {
			return item
		}
	}
	return model.Currency{}
}
func (bot *tgbot) getFloat(uc *UserChat, messageChannel chan string, text string, parentCtx context.Context) (float32, error) {
	bot.sendText(uc.chatId, text)

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.RegisterChannel(uc, "message", "", messageChannel, ctx)

	return getFloatFromChannel(messageChannel, ctx)
}
func (bot *tgbot) getString(uc *UserChat, messageChannel chan string, text string, parentCtx context.Context) (string, error) {
	bot.sendText(uc.chatId, text)

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.RegisterChannel(uc, "message", "", messageChannel, ctx)
	val, err := getStringFromChannel(messageChannel, ctx)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (bot *tgbot) RegisterChannel(uc *UserChat, mediaType string, callbackInfo string, data chan<- string, ctx context.Context) {
	messageChannel, err := bot.AdvancedMode().RegisterChannel(strconv.Itoa(uc.chatId), mediaType)
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
					if up.Message.From.Id == uc.userId {
						data <- up.Message.Text
					}
				case "callback_query":
					ar = strings.Split(up.CallbackQuery.Data, ": ")
					if up.CallbackQuery.From.Id == uc.userId && ar[0] == callbackInfo {
						data <- ar[1]
					}
				}
			}
		case <-ctx.Done():
			{
				bot.AdvancedMode().UnRegisterChannel(strconv.Itoa(uc.chatId), mediaType)
				return
			}
		}
	}
}

func (bot *tgbot) cancelOperation(u *objects.Update) {
	uc := UserChat{u.Message.From.Id, u.Message.Chat.Id}
	cancel, ok := bot.taskCancel.Load(uc)
	if ok {
		cancel()
	}
}

func (bot tgbot) sendInlineKeyboard(uc *UserChat, text string, kb telego.MarkUps) (*objects.Message, error) {
	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, text, "", 0, false, false, nil, false, false, kb)
	if err != nil {
		return nil, err
	}
	return msg.Result, nil

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
