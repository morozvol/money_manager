package tgbot

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
)

func (bot *tgbot) categoriesKeyboard(uc *o.UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context) (*model.Category, error) {
	categories := bot.getUserCategories(uc.UserId)
	id := 1
	lastId := 1

	msg, err := bot.sendInlineKeyboard(uc, "Выбор категории", nil)
	if err != nil {
		bot.error(err, "categoriesKeyboard: "+ErrSendMessage.Error(), nil)
	}

	defer func() {
		_, err := editor.DeleteMessage(msg.MessageId)
		if err != nil {
			bot.error(err, "categoriesKeyboard: "+ErrDeleteMessage.Error(), msg)
		}
	}()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.registerChannel(uc, "callback_query", "id category", messageChannel, ctx)

	for {
		kb := bot.CreateInlineKeyboard()

		cat := categories.GetCategoriesByIdParent(id)

		for i, c := range cat {
			kb.AddCallbackButton(fmt.Sprintf("%s", c.Name), fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
		}
		if id != 1 {
			kb.AddCallbackButton(fmt.Sprintf("%s", "Назад"), fmt.Sprintf("id category: %d", lastId), int((len(cat)-1)/2)+1+1)
		}
		_, err = editor.EditReplyMarkup(msg.MessageId, "", kb)
		if err != nil {
			bot.error(err, "categoriesKeyboard: "+ErrEditMessage.Error(), msg)
		}

		val, err := getIntFromChannel(messageChannel, ctx)
		if err != nil {
			return nil, err
		}
		lastId = id
		id = val

		if id != 1 {
			category, err := categories.GetCategoryById(id)
			if err != nil {
				bot.error(err, "categoriesKeyboard: Категория не существует", id)
				return nil, ErrUnknown
			}

			if category.IsEnd {
				return category, nil
			}
		}
	}
}

func (bot *tgbot) getUserCategories(userId int) model.Categories {
	category, err := bot.store.Category().Get(userId)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
	return category
}
