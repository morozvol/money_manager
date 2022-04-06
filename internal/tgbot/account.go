package tgbot

import (
	"bytes"
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
)

func (bot *tgbot) addAccount(u *objects.Update) {
	chatId := u.Message.Chat.Id
	userId := u.Message.From.Id

	_, err := bot.store.User().Find(userId)
	if err != nil {
		bot.help(u)
		return
	}
	uc := &userChat{userId, chatId}
	account := model.Account{}
	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.chatId)

	account.Name = bot.getString(uc, msgChannel, "Введите название кошелька")
	account.Currency = *bot.currencyKeyboard(uc, msgChannel, msgEditor)
	account.Balance = bot.getFloat(uc, msgChannel, "Введите сумму")
	account.IdUser = uc.userId

	err = bot.store.Account().Create(&account)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
}

func (bot *tgbot) accountsKeyboard(uc *userChat, messageChannel chan string, editor *bt.MessageEditor) *model.Account {
	accounts := bot.getUserAccounts(uc.userId)
	if len(accounts) == 0 {
		bot.youNeedCreateAccount(uc)
		bot.Logger.Info("accountsKeyboard: Отправлено сообщение пользователю о необходимости создать счёт")
		return nil
	}
	kb := bot.CreateInlineKeyboard()
	for i, account := range accounts {
		kb.AddCallbackButton(fmt.Sprintf("%s %s: %.2f", account.Name, account.Currency.Code, account.Balance),
			fmt.Sprintf("id account: %d", account.Id), i+1)
	}

	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, "Выбор счёта", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.error(err, "accountsKeyboard", nil)
	}

	defer func() {
		_, err := editor.DeleteMessage(msg.Result.MessageId)
		if err != nil {
			bot.error(err, "accountsKeyboard: не удалось удалить сообщение", msg)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id account", messageChannel, ctx)

	if val, err := strconv.ParseInt(<-messageChannel, 10, 64); err == nil {
		for _, a := range accounts {
			if a.Id == val {
				return &a
			}
		}
	}
	return nil
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

func (bot *tgbot) getAccountsInfo(u *objects.Update) {
	accounts := bot.getUserAccounts(u.Message.From.Id)
	var buffer bytes.Buffer
	for _, account := range accounts {
		buffer.WriteString(fmt.Sprintf("  %s %s: %.2f\n", account.Name, account.Currency.Code, account.Balance))
	}
	bot.sendText(u.Message.Chat.Id, buffer.String())
}

func (bot *tgbot) youNeedCreateAccount(uc *userChat) {
	bot.sendText(uc.chatId, "Вам необходимо создать счёт перед добавлением операции. Это можно сделать при помощи /add_account")
}
