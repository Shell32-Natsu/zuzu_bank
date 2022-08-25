package commands

import (
	"context"
	"fmt"
	"github.com/Shell32-Natsu/zuzu_bank/internal/logging"
	"strings"

	"github.com/Shell32-Natsu/zuzu_bank/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Balance struct{}

func (b Balance) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := &strings.Builder{}

	queryId := b.QueryUserId(msg)
	if queryId < 0 {
		return tgbotapi.MessageConfig{}, fmt.Errorf("invalid user id")
	}
	logging.LogDebugf("user id to query: %d", queryId)

	balance, err := queryBalance(ctx, queryId)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to query balance: %s", err)
	}
	writeStringf(text, "User ID:%d\n", queryId)
	if queryId == msg.From.ID {
		// Query is for user self
		writeStringf(text, "User name: %s\n", msg.From.UserName)
	}
	writeStringf(text, "当前余额：%d\n", balance)

	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func queryBalance(ctx context.Context, id int64) (int64, error) {
	return db.Balance(ctx, id)
}

func (Balance) Help() string {
	return "用户的猪猪币余额\n\n用法：/balance [user_id]\n\n如果不提供user_id则查询当前用户，只有银行员工可以查询其他用户"
}

func (Balance) QueryUserId(msg *tgbotapi.Message) int64 {
	const userIdPos = 0
	return getQueryUserIdFromCommandArgs(msg.From.ID, msg.CommandArguments(), userIdPos)
}
