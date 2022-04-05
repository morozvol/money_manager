package tgbot

import (
	"fmt"
	"github.com/morozvol/money_manager/internal/config"
	"github.com/morozvol/money_manager/internal/store"
	"github.com/morozvol/money_manager/internal/store/sqlstore"
	"github.com/morozvol/money_manager/internal/store/sqlstore/db"
	bt "github.com/morozvol/telego"
	cfg "github.com/morozvol/telego/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type userChat struct {
	userId int
	chatId int
}

type tgbot struct {
	bt.Bot
	Logger   *zap.Logger
	store    store.Store
	userData map[userChat]chan string
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

	return &tgbot{*bot, zap.L(), s, make(map[userChat]chan string)}, nil
}

func (bot *tgbot) HandlersRegister() error {
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

	messageChannel, err := bot.AdvancedMode().RegisterChannel("", "")
	if err != nil {
		bot.Logger.Fatal(err.Error())
	}

	for {
		up := <-*messageChannel
		switch up.GetType() {
		case "message":
			{
				fmt.Println(up.Message.Text)
				bot.sendData(userChat{up.Message.From.Id, up.Message.Chat.Id}, up.Message.Text)
			}
		case "callback_query":
			uc := userChat{up.CallbackQuery.From.Id, up.CallbackQuery.Message.Chat.Id}
			ar := strings.Split(up.CallbackQuery.Data, ": ")
			switch ar[0] {
			case "id account":
				bot.sendData(uc, ar[1])
			case "id currency":
				bot.sendData(uc, ar[1])
			case "id category":
				bot.sendData(uc, ar[1])
			}
		default:
			continue
		}
	}
}
