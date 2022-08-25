package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Shell32-Natsu/zuzu_bank/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Deposit struct{}

func (d Deposit) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	var e tgbotapi.MessageConfig
	text := &strings.Builder{}

	queryId := d.QueryUserId(msg)
	if queryId < 0 {
		return e, fmt.Errorf("invalid user id")
	}

	n, err := number(msg)
	if err != nil {
		return e, fmt.Errorf("failed to get deposit number: %s", err)
	}

	desc, err := descption(msg)
	if err != nil {
		return e, fmt.Errorf("failed to get deposit description: %s", desc)
	}

	old, new, err := db.Deposit(ctx, queryId, n, desc)
	if err != nil {
		return e, fmt.Errorf("failed to add deposit: %s", err)
	}

	writeStringf(text, "User ID:%d\n", queryId)
	if queryId == msg.From.ID {
		// Query is for user self
		writeStringf(text, "User name: %s\n", msg.From.UserName)
	}

	writeStringf(text, "原余额：%d\n", old)
	writeStringf(text, "存入：%d\n", n)
	writeStringf(text, "新余额：%d\n", new)

	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func number(msg *tgbotapi.Message) (int64, error) {
	args := strings.Split(msg.CommandArguments(), " ")
	if len(args) < 1 {
		return 0, fmt.Errorf("deposit numebr is required")
	}
	return strconv.ParseInt(args[0], 10, 0)
}

func descption(msg *tgbotapi.Message) (string, error) {
	args := strings.Split(msg.CommandArguments(), " ")
	if len(args) < 2 {
		return "", fmt.Errorf("description is required")
	}
	return args[1], nil
}

func (Deposit) Help() string {
	return "存猪猪币\n\n用法：/deposit number description [user_id]"
}

func (Deposit) QueryUserId(msg *tgbotapi.Message) int64 {
	const userIdPos = 2
	return getQueryUserIdFromCommandArgs(msg.From.ID, msg.CommandArguments(), userIdPos)
}
