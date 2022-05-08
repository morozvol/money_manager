package main

import (
	"github.com/morozvol/money_manager/internal/config"
	"github.com/morozvol/money_manager/internal/tgbot"
	"github.com/morozvol/money_manager/pkg/store/sqlstore"
	"github.com/morozvol/money_manager/pkg/store/sqlstore/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	conf, err := config.Init()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger, _ = zap.NewProduction()
	logger.WithOptions(zap.IncreaseLevel(zapcore.DebugLevel))
	zap.ReplaceGlobals(logger)

	dataBase, err := db.New(conf.DB, logger)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	store := sqlstore.New(dataBase)
	bot, err := tgbot.New(store, conf)
	if err != nil {
		println(err)
		return
	}
	bot.Logger.Fatal(bot.Start().Error())

}
