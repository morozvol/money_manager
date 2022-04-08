package tgbot

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"
	"github.com/morozvol/money_manager/internal/model"
	"html"
)

func (bot *tgbot) addAccount(u *objects.Update) {
	uc := &UserChat{u.Message.From.Id, u.Message.Chat.Id}

	ctx, cancel := context.WithCancel(context.Background())
	bot.taskCancel.Store(*uc, cancel)

	user, err := bot.store.User().Find(uc.userId)
	if err != nil {
		bot.help(u)
		return
	}
	account := model.Account{}
	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.chatId)

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

	account.IdUser = uc.userId

	err = bot.store.Account().Create(&account)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
}

func (bot *tgbot) accountsKeyboard(uc *UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context) (*model.Account, error) {
	accounts := bot.getUserAccounts(uc.userId)
	if len(accounts) == 0 {
		bot.youNeedCreateAccount(uc)
		bot.Logger.Info("accountsKeyboard: Отправлено сообщение пользователю о необходимости создать счёт")
		return nil, errors.New("Отправлено сообщение пользователю о необходимости создать счёт")
	}
	kb := bot.CreateInlineKeyboard()
	for i, account := range accounts {
		kb.AddCallbackButton(fmt.Sprintf("%s\t    %s %s: %.2f", html.UnescapeString(account.AccountType.Symbol), account.Name, account.Currency.Code, account.Balance),
			fmt.Sprintf("id account: %d", account.Id), i+1)
	}

	msg, err := bot.sendInlineKeyboard(uc, "Выбор счёта", kb)
	if err != nil {
		bot.error(err, "accountsKeyboard", nil)
	}

	defer func() {
		_, err := editor.DeleteMessage(msg.MessageId)
		if err != nil {
			bot.error(err, "accountsKeyboard: не удалось удалить сообщение", msg)
		}
	}()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id account", messageChannel, ctx)

	val, err := getIntFromChannel(messageChannel, ctx)
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		if a.Id == int64(val) {
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
	var buffer bytes.Buffer
	for _, account := range accounts {
		buffer.WriteString(fmt.Sprintf("  %s %s: %.2f\n", account.Name, account.Currency.Code, account.Balance))
	}
	bot.sendText(u.Message.Chat.Id, buffer.String())
}

func (bot *tgbot) youNeedCreateAccount(uc *UserChat) {
	bot.sendText(uc.chatId, "Вам необходимо создать счёт перед добавлением операции. Это можно сделать при помощи /add_account")
}
