package tgbot

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/model/category_tree"
)

func (bot *tgbot) categoriesKeyboard(uc *o.UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context) (*model.Category, error) {
	categories := bot.getUserCategories(uc.UserId)
	OpenNode := categories.Root

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

		cat := OpenNode.GetChildren().ToCategories()

		for i, c := range cat {
			kb.AddCallbackButton(fmt.Sprintf("%s", c.Name), fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
		}
		if OpenNode != categories.Root {
			kb.AddCallbackButton(fmt.Sprintf("%s", "Назад"), fmt.Sprintf("id category: %d", OpenNode.Category.IdParent), int((len(cat)-1)/2)+1+1)
		}
		_, err = editor.EditReplyMarkup(msg.MessageId, "", kb)
		if err != nil {
			bot.error(err, "categoriesKeyboard: "+ErrEditMessage.Error(), msg)
		}

		id, err := getIntFromChannel(messageChannel, ctx)
		if err != nil {
			return nil, err
		}

		OpenNode, _ = categories.FindNode(id)

		category := OpenNode.Category

		if category.IsEnd {
			return category, nil
		}
	}
}

func (bot *tgbot) getUserCategories(userId int) *category_tree.CategoryTree {
	category, err := bot.store.Category().Get(userId)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
	return category
}
