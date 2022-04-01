package main

import (
	"github.com/morozvol/money_manager/internal/tgbot"
)

func main() {
	bot, err := tgbot.New()
	if err != nil {
		println(err)
		return
	}
	bot.Logger.Fatal(bot.HandlersRegister().Error())

}
