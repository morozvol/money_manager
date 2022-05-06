package tgbot

import (
	"context"
	"fmt"
	bt "github.com/SakoDroid/telego"
	objs "github.com/SakoDroid/telego/objects"
	o "github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/model"
	"html"
)

const (
	edit = iota
	deleteCategory
	createCategory
	createDirectory
	open
	save
)
const (
	folderSymbol      = "&#128193;"
	endCategorySymbol = "&#128204;"
	deleteSymbol      = "&#9940;"
	editSymbol        = "&#128221;"
)

func (bot *tgbot) editCategories(uc *o.UserChat, messageChannel chan string, editor *bt.MessageEditor, parentCtx context.Context) error {
	categories := bot.getUserCategories(uc.UserId)
	var symbol string
	isUpdateKeyboard := true
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

	go bot.registerChannel(uc, "callback_query", "id category editor", messageChannel, ctx)

	for {
		if isUpdateKeyboard {
			kb := bot.CreateInlineKeyboard()

			cat := categories.GetCategoriesByIdParent(id)
			line := 1
			for _, c := range cat {
				if c.IsEnd {
					symbol = endCategorySymbol
				} else {
					symbol = folderSymbol
				}
				kb.AddCallbackButton(fmt.Sprintf("%s %s", html.UnescapeString(symbol), c.Name), fmt.Sprintf("id category editor: %d %d", c.Id, open), line)

				kb.AddCallbackButton(fmt.Sprintf("%s %s", html.UnescapeString(editSymbol), "Изменить"), fmt.Sprintf("id category editor: %d %d", c.Id, edit), line)
				kb.AddCallbackButton(fmt.Sprintf("%s %s", html.UnescapeString(deleteSymbol), "Удалить"), fmt.Sprintf("id category editor: %d %d", c.Id, deleteCategory), line)
				line++
			}

			kb.AddCallbackButton(fmt.Sprintf("%s", "Создать папку"), fmt.Sprintf("id category editor: %d %d", id, createDirectory), line)
			kb.AddCallbackButton(fmt.Sprintf("%s", "Создать категорию"), fmt.Sprintf("id category editor: %d %d", id, createCategory), line)
			line++
			if id != 1 {
				kb.AddCallbackButton(fmt.Sprintf("%s", " <- Назад"), fmt.Sprintf("id category editor: %d %d", lastId, open), line)
			}
			kb.AddCallbackButton(fmt.Sprintf("%s", "Сохранить изменения"), fmt.Sprintf("id category editor: %d %d", 1, save), line)

			_, err = editor.EditReplyMarkup(msg.MessageId, "", kb)
			if err != nil {
				bot.error(err, "categoriesKeyboard: "+ErrEditMessage.Error(), msg)
			}
		}

		idCategory, action, err := getTwoIntFromChannel(messageChannel, ctx)
		if err != nil {
			return err
		}
		category, err := categories.GetCategoryById(idCategory)
		if err != nil {
			bot.error(err, "categoriesKeyboard: Категория не существует", id)
			return ErrUnknown
		}

		switch action {
		case open:
			if category.IsEnd {
				isUpdateKeyboard = false
			} else {
				lastId = id
				id = idCategory
				isUpdateKeyboard = true
			}

		case deleteCategory:
			categories = categories.DeleteById(idCategory)
			isUpdateKeyboard = true
		case edit:
			//edit category with id = idCategory
			isUpdateKeyboard = true
		case createCategory:
			cat := &model.Category{IdOwner: uc.UserId, IdParent: idCategory, IsEnd: true, IsSystem: false}
			err := bot.createCategory(cat)
			if err != nil {
				return err
			}
			categories = append(categories, *cat)
			isUpdateKeyboard = true
		case createDirectory:
			cat := &model.Category{IdOwner: uc.UserId, IdParent: idCategory, IsEnd: false, IsSystem: false}
			err := bot.createCategory(cat)
			if err != nil {
				return err
			}
			categories = append(categories, *cat)
			isUpdateKeyboard = true
		case save:
			return nil
		}
	}
}

func (bot *tgbot) editCategory(u *objs.Update) {
	uc := &o.UserChat{UserId: u.Message.From.Id, ChatId: u.Message.Chat.Id}
	if !bot.beforeExecution(uc) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	bot.taskCancel.Store(*uc, cancel)
	defer bot.taskCancel.Cancel(*uc)

	msgChannel := make(chan string)
	defer close(msgChannel)
	msgEditor := bot.GetMsgEditor(uc.ChatId)

	err := bot.editCategories(uc, msgChannel, msgEditor, ctx)
	if err != nil {
		return
	}
}

func (bot *tgbot) createCategory(cat *model.Category) error {
	cat.Name = "test Category" //get
	if !cat.IsEnd {
		cat.Type = model.Consumption //default value
	}
	cat.Type = model.Coming //get

	err := bot.store.Category().Create(cat)
	if err != nil {
		return err
	}
	return nil
}

func (bot *tgbot) deleteCategory(id int) error {
	err := bot.store.Category().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
