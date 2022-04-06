package tgbot

import (
	"context"
	"fmt"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
	"strings"
)

func (bot *tgbot) sendText(chatId int, text string) {
	_, err := bot.SendMessage(chatId, text, "", 0, false, false)
	if err != nil {
		bot.error(err, "sendText: Не удалось отправить сообщение", nil)
	}
}

func (bot *tgbot) error(err error, msg string, data interface{}) {
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
func (bot *tgbot) getFloat(uc *userChat, messageChannel chan string, text string) float32 {
	bot.sendText(uc.chatId, text)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go bot.RegisterChannel(uc, "message", "", messageChannel, ctx)

	s := <-messageChannel
	s = strings.Replace(s, ",", ".", -1)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		bot.error(err, "getFloat: Не удалось привести к float", nil)
		return 0
	}
	return float32(val)
}
func (bot *tgbot) getString(uc *userChat, messageChannel chan string, text string) string {
	bot.sendText(uc.chatId, text)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go bot.RegisterChannel(uc, "message", "", messageChannel, ctx)

	return <-messageChannel

}

func (bot *tgbot) RegisterChannel(uc *userChat, mediaType string, callbackInfo string, data chan<- string, ctx context.Context) {
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
