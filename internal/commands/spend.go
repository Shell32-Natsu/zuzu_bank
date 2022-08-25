package commands

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Spend struct{}

func (Spend) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}

	text.WriteString(fmt.Sprintf("User ID:%d\nUser Name:%s\n", msg.From.ID, msg.From.UserName))
	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func (Spend) Help() string {
	return "花猪猪币"
}

func (Spend) QueryUserId(msg *tgbotapi.Message) int64 {
	return 0
}
