package commands

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var commandMap = map[string]BotCommand{
	"help":        Help{},
	"balance":     Balance{},
	"deposit":     Deposit{},
	"spend":       Spend{},
	"transaction": Transaction{},
}

type BotCommand interface {
	Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	Help() string
	QueryUserId(msg *tgbotapi.Message) int64
}
