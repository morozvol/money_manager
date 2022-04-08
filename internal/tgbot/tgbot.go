package tgbot

import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	"github.com/morozvol/money_manager/internal/config"
	"github.com/morozvol/money_manager/internal/store"
	"github.com/morozvol/money_manager/internal/store/sqlstore"
	"github.com/morozvol/money_manager/internal/store/sqlstore/db"
	"github.com/morozvol/money_manager/internal/tgbot/objects"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type tgbot struct {
	bt.Bot
	Logger     *zap.Logger
	store      store.Store
	taskCancel objects.CancelFuncMap
}

func New() (*tgbot, error) {
	var err error

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	conf, err := config.Init()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger, _ = zap.NewDevelopment()
	logger.WithOptions(zap.IncreaseLevel(zapcore.DebugLevel))
	zap.ReplaceGlobals(logger)

	if err := os.Setenv("API_KEY", conf.ApiKey); err != nil {
		return nil, err
	}

	up := cfg.DefaultUpdateConfigs()

	cf := cfg.BotConfigs{
		BotAPI:         cfg.DefaultBotAPI,
		APIKey:         os.Getenv("API_KEY"),
		UpdateConfigs:  up,
		Webhook:        false,
		LogFileAddress: cfg.DefaultLogFile,
	}

	var bot *bt.Bot

	bot, err = bt.NewBot(&cf)
	if err != nil {
		return nil, err
	}

	if err := bot.Run(); err != nil {
		return nil, err
	}

	dataBase, err := db.New(conf, logger)
	if err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	s := sqlstore.New(dataBase)

	return &tgbot{*bot, zap.L(), s, objects.NewCancelFuncMap()}, nil
}

func (bot *tgbot) Start() error {
	if err := bot.AddHandler("/register", bot.register, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/help", bot.help, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/add_account", bot.addAccount, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/add_operation", bot.addOperation, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/info", bot.getAccountsInfo, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/cancel", bot.cancelOperation, "private", "group"); err != nil {
		return err
	}

	messageChannel := bot.GetUpdateChannel()

	defer bot.Stop()

	for {
		<-*messageChannel
	}
}
