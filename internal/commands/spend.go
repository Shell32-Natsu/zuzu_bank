package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shell32-Natsu/zuzu_bank/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Spend struct{}

func (s Spend) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	var e tgbotapi.MessageConfig
	text := &strings.Builder{}

	queryId := s.QueryUserId(msg)
	if queryId < 0 {
		return e, fmt.Errorf("invalid user id")
	}

	n, err := number(msg)
	if err != nil {
		return e, fmt.Errorf("failed to get spend number: %s", err)
	}

	desc, err := descption(msg)
	if err != nil {
		return e, fmt.Errorf("failed to get spend description: %s", desc)
	}

	old, new, err := db.Spend(ctx, queryId, n, desc)
	if err != nil {
		return e, fmt.Errorf("failed to add spend: %s", err)
	}

	writeStringf(text, "User ID:%d\n", queryId)
	if queryId == msg.From.ID {
		// Query is for user self
		writeStringf(text, "User name: %s\n", msg.From.UserName)
	}

	writeStringf(text, "原余额：%d\n", old)
	writeStringf(text, "花费：%d\n", n)
	writeStringf(text, "新余额：%d\n", new)

	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func (Spend) Help() string {
	return "花猪猪币\n\n用法：/spend number description [user_id]"
}

func (Spend) QueryUserId(msg *tgbotapi.Message) int64 {
	const userIdPos = 2
	return getQueryUserIdFromCommandArgs(msg.From.ID, msg.CommandArguments(), userIdPos)
}
