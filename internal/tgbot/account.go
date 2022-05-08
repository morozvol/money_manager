package tgbot

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"html"
)

func (bot *tgbot) addAccount(u *objects.Update) {
	uc := &o.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}
	if !bot.beforeExecution(uc) {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	bot.taskCancel.Store(*uc, cancel)
	defer bot.taskCancel.Cancel(*uc)

	user, err := bot.store.User().Find(uc.UserId)
	if err != nil {
		return
	}
	account := model.Account{}
	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.ChatId)

	name, err := bot.getString(uc, msgChannel, "Введите название кошелька", ctx)
	if err != nil {
		return
	}
	account.Name = name

	currency, err := bot.chooseCurrency(user, uc, msgChannel, msgEditor, ctx, false)
	if err != nil {
		return
	}
	account.Currency = *currency

	bal, err := bot.getFloat(uc, msgChannel, "Введите сумму", ctx)
	if err != nil {
		return
	}
	account.Balance = bal

	account.IdUser = uc.UserId

	err = bot.store.Account().Create(&account)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
}

func (bot *tgbot) accountsKeyboard(uc *o.UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context) (*model.Account, error) {
	accounts := bot.getUserAccounts(uc.UserId)
	if len(accounts) == 0 {
		bot.youNeedCreateAccount(uc)
		bot.Logger.Info("accountsKeyboard: Отправлено сообщение пользователю о необходимости создать счёт")
		return nil, errors.New("отправлено сообщение пользователю о необходимости создать счёт")
	}
	kb := bot.CreateInlineKeyboard()
	for i, account := range accounts {
		kb.AddCallbackButton(fmt.Sprintf("%s\t    %s %s: %.2f", html.UnescapeString(account.AccountType.Symbol), account.Name, account.Currency.Code, account.Balance),
			fmt.Sprintf("id account: %d", account.Id), i+1)
	}

	msg, err := bot.sendInlineKeyboard(uc, "Выбор счёта", kb)
	if err != nil {
		bot.error(err, "accountsKeyboard: "+ErrSendMessage.Error(), nil)
	}

	defer func() {
		_, err := editor.DeleteMessage(msg.MessageId)
		if err != nil {
			bot.error(err, "accountsKeyboard: "+ErrDeleteMessage.Error(), msg)
		}
	}()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.registerChannel(uc, "callback_query", "id account", messageChannel, ctx)

	val, err := getIntFromChannel(messageChannel, ctx)
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		if a.Id == val {
			return &a, nil
		}
	}
	return nil, ErrUnknown
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
	if len(accounts) == 0 {
		bot.sendText(u.Message.Chat.Id, "Нет счетов для вывода информации")
		return
	}
	var buffer bytes.Buffer
	for _, account := range accounts {
		buffer.WriteString(fmt.Sprintf("  %s %s: %.2f\n", account.Name, account.Currency.Code, account.Balance))
	}
	bot.sendText(u.Message.Chat.Id, buffer.String())
}

func (bot *tgbot) youNeedCreateAccount(uc *o.UserChat) {
	bot.sendText(uc.ChatId, "Вам необходимо создать счёт перед добавлением операции. Это можно сделать при помощи /add_account")
}
