package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Balance struct{}

func (mb Balance) Run(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}

	text.WriteString(fmt.Sprintf("User ID:%d\nUser Name:%s\n", msg.From.ID, msg.From.UserName))
	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func (mb Balance) Help() string {
	return "Show your current balance in account"
}
