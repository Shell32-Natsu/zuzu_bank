package commands

import (
	"fmt"
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
	Run(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error)
	Help() string
}

func RunCommand(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	for c, h := range commandMap {
		if c == msg.Command() {
			return h.Run(msg)
		}
	}
	// Unknown command
	return tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("命令不存在: %s", msg.Command())), nil
}
