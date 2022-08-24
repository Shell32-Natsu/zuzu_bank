package commands

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Deposit struct{}

func (Deposit) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}

	text.WriteString(fmt.Sprintf("User ID:%d\nUser Name:%s\n", msg.From.ID, msg.From.UserName))
	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func (Deposit) Help() string {
	return "存猪猪币"
}
