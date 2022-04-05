package tgbot

import (
	"fmt"
	bt "github.com/SakoDroid/telego"
	"github.com/morozvol/money_manager/internal/model"
	"strconv"
)

func (bot *tgbot) categoriesKeyboard(uc *userChat) *model.Category {
	categories := bot.getUserCategories(uc.userId)
	id := 0
	lastId := 0

	kb := bot.CreateInlineKeyboard()
	for i, c := range getCategoriesByIdParent(id, &categories) {
		kb.AddCallbackButton(fmt.Sprintf("%s", c.Name),
			fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
	}

	msg, err := bot.AdvancedMode().ASendMessage(uc.chatId, "Выбор категории", "", 0, false, false, nil, false, false, kb)
	if err != nil {
		bot.Error(err, "categoriesKeyboard: ообщение не отправлено", nil)
	}
	msgEditor := bot.GetMsgEditor(uc.chatId)

	defer func(msgEditor *bt.MessageEditor, messageId int) {
		_, err := msgEditor.DeleteMessage(messageId)
		if err != nil {
			bot.Error(err, "categoriesKeyboard: не удалось удалить сообщение", msg)
		}
	}(msgEditor, msg.Result.MessageId)

	for {
		if val, err := strconv.ParseInt(<-bot.userData[*uc], 10, 64); err == nil {
			lastId = id
			id = int(val)
		} else {
			bot.Error(err, "categoriesKeyboard: Не удалось понять что выбрал пользователь", nil)
		}
		if id != 0 {
			category, err := getCategoryById(int64(id), &categories)
			if err != nil {
				bot.Error(err, "categoriesKeyboard: Категория не существует", id)
				return nil
			}

			if category.IsEnd {
				return category
			}
		}
		kb := bot.CreateInlineKeyboard()

		cat := getCategoriesByIdParent(id, &categories)

		for i, c := range cat {
			kb.AddCallbackButton(fmt.Sprintf("%s", c.Name), fmt.Sprintf("id category: %d", c.Id), int(i/2)+1)
		}
		if id != 0 {
			kb.AddCallbackButton(fmt.Sprintf("%s", "Назад"), fmt.Sprintf("id category: %d", lastId), int((len(cat)-1)/2)+1+1)

		}
		_, err = msgEditor.EditReplyMarkup(msg.Result.MessageId, "", kb)
		if err != nil {
			bot.Error(err, "categoriesKeyboard: Не удалось изменить сообщение", msg)
		}
	}
}

func (bot *tgbot) getUserCategories(userId int) []model.Category {
	category, err := bot.store.Category().GetAll(userId)
	if err != nil {
		bot.Logger.Error(err.Error())
	}
	return category
}

func getCategoriesByIdParent(id int, categories *[]model.Category) []model.Category {
	var res []model.Category

	for _, c := range *categories {
		if c.IdParent == id {
			res = append(res, c)
		}
	}
	return res
}

func getCategoryById(id int64, categories *[]model.Category) (*model.Category, error) {

	for _, c := range *categories {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("Category with id=%d does not exist ", id)
}
