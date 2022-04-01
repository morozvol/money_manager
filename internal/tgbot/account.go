package tgbot

import (
	"bytes"
	"fmt"
	"github.com/morozvol/money_manager/internal/model"
	objs "github.com/morozvol/telego/objects"
	"strconv"
)

func (bot *tgbot) addAccount(u *objs.Update) {
	chatId := u.Message.Chat.Id
	userId := u.Message.From.Id

	_, err := bot.store.User().Find(userId)
	if err != nil {
		bot.help(u)
		return
	}
	uc := userChat{userId, chatId}
	account := model.Account{}
	bot.userData[uc] = make(chan string)
	bot.SendText(chatId, "Введите название кошелька")
	account.Name = <-bot.userData[uc]
	bot.currencyKeyboard(u)
	if val, err := strconv.ParseInt(<-bot.userData[uc], 10, 64); err == nil {
		account.Currency.Id = val
	}
	bot.SendText(chatId, "Введите сумму")
	if fval, err := strconv.ParseFloat(<-bot.userData[uc], 64); err == nil {
		account.Balance = float32(fval)
	}

	close(bot.userData[uc])
	delete(bot.userData, uc)

	account.IdUser = uc.userId

	err = bot.store.Account().Create(&account)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
}

func (bot *tgbot) accountsKeyboard(u *objs.Update) {
	accounts := bot.getUserAccounts(u.Message.From.Id)
	kb := bot.CreateInlineKeyboard()
	for i, account := range accounts {
		kb.AddCallbackButton(fmt.Sprintf("%s %s: %.2f", account.Name, account.Currency.Code, account.Balance),
			fmt.Sprintf("id account: %d", account.Id), i+1)
	}

	_, err := bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "Выбор счёта", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		fmt.Println(err)
	}
}

func (bot *tgbot) getUserAccounts(userId int) []model.Account {
	currencies, err := bot.store.Currency().GetAll()
	if err != nil {
		bot.Logger.Error(err.Error())
	}

	accounts, err := bot.store.Account().FindByUserId(userId)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
	for i := range accounts {
		accounts[i].Currency = getCurrencyById(accounts[i].Currency.Id, currencies)
	}

	return accounts
}

func (bot *tgbot) getAccountsInfo(u *objs.Update) {
	accounts := bot.getUserAccounts(u.Message.From.Id)
	var buffer bytes.Buffer
	for _, account := range accounts {
		buffer.WriteString(fmt.Sprintf("  %s %s: %.2f\n", account.Name, account.Currency.Code, account.Balance))
	}
	bot.SendText(u.Message.Chat.Id, buffer.String())
}
