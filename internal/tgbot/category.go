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
	id := 0
	lastId := 0

	kb := bot.CreateInlineKeyboard()
	for i, c := range categories.GetCategoriesByIdParent(id) {
		kb.AddCallbackButton(fmt.Sprintf("%s", c.Name),
			fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
	}

	msg, err := bot.sendInlineKeyboard(uc, "Выбор категории", kb)
	if err != nil {
		bot.error(err, "categoriesKeyboard: ообщение не отправлено", nil)
	}

	defer func() {
		_, err := editor.DeleteMessage(msg.MessageId)
		if err != nil {
			bot.error(err, "categoriesKeyboard: не удалось удалить сообщение", msg)
		}
	}()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	go bot.RegisterChannel(uc, "callback_query", "id category", messageChannel, ctx)

	for {
		val, err := getIntFromChannel(messageChannel, ctx)
		if err != nil {
			return nil, err
		}
		lastId = id
		id = val

		if id != 0 {
			category, err := categories.GetCategoryById(int64(id))
			if err != nil {
				bot.error(err, "categoriesKeyboard: Категория не существует", id)
				return nil, ErrUnknown
			}

			if category.IsEnd {
				return category, nil
			}
		}
		kb := bot.CreateInlineKeyboard()

		cat := categories.GetCategoriesByIdParent(id)

		for i, c := range cat {
			kb.AddCallbackButton(fmt.Sprintf("%s", c.Name), fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
		}
		if id != 0 {
			kb.AddCallbackButton(fmt.Sprintf("%s", "Назад"), fmt.Sprintf("id category: %d", lastId), int((len(cat)-1)/2)+1+1)

		}
		_, err = editor.EditReplyMarkup(msg.MessageId, "", kb)
		if err != nil {
			bot.error(err, "categoriesKeyboard: Не удалось изменить сообщение", msg)
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
