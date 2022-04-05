package tgbot

import (
	"fmt"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
	"strings"
)

func (bot *tgbot) SendText(chatId int, text string) {
	_, err := bot.SendMessage(chatId, text, "", 0, false, false)
	if err != nil {
		bot.Error(err, "SendText: Не удалось отправить сообщение", nil)
	}
}

func (bot *tgbot) sendData(uc userChat, data string) {
	if _, ok := bot.userData[uc]; !ok {
		return
	}
	bot.userData[uc] <- data
}
func (bot *tgbot) Error(err error, msg string, data interface{}) {
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
func (bot *tgbot) getFloat(uc *userChat, text string) float32 {
	bot.SendText(uc.chatId, text)
	s := <-bot.userData[*uc]
	s = strings.Replace(s, ",", ".", -1)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		bot.Error(err, "getFloat: Не удалось привести к float", nil)
		return 0
	}
	return float32(val)
}
