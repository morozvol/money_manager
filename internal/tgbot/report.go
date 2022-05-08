package tgbot

import (
	objs "github.com/SakoDroid/telego/objects"
	"time"
)

func (bot *tgbot) getSpendingPerWeek(u *objs.Update) {
	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	day := int(today.Weekday()) - 1
	if day < 0 {
		day = 6 //первй день недели понедельник вместо воскресенья
	}
	periodFrom := today.AddDate(0, 0, -day)
	periodTo := today.AddDate(0, 0, 1)

	_, err := bot.store.Operation().Get(periodFrom, periodTo, u.Message.From.Id)
	if err != nil {
		return
	}
}
