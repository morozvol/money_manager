package tgbot

import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	"github.com/morozvol/money_manager/internal/config"
	"github.com/morozvol/money_manager/internal/tgbot/objects"
	"github.com/morozvol/money_manager/pkg/store"
	"go.uber.org/zap"
	"os"
)

type tgbot struct {
	bt.Bot
	Logger     *zap.Logger
	store      store.Store
	taskCancel objects.CancelFuncMap
}

func New(store store.Store, conf *config.Config) (*tgbot, error) {
	var err error

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

	return &tgbot{*bot, zap.L(), store, objects.NewCancelFuncMap()}, nil
}

func (bot *tgbot) Start() error {

	if err := bot.AddHandler("/help", bot.help, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/add_account", bot.addAccount, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/add_operation", bot.addPaymentOperation, "private", "group"); err != nil {
		return err
	}
	if err := bot.AddHandler("/add_transfer", bot.addTransferOperation, "private", "group"); err != nil {
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
