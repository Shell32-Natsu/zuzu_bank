package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	"github.com/Shell32-Natsu/zuzu_bank/internal/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Balance struct{}

func (Balance) Run(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}

	text.WriteString(fmt.Sprintf("User ID:%d\nUser Name:%s\n", msg.From.ID, msg.From.UserName))

	queryId := msg.From.ID
	args := msg.CommandArguments()
	if args != "" {
		if !config.IsAdmin(msg.From.ID) {
			text.WriteString("只有银行员工可以查询其他账户")
			return tgbotapi.NewMessage(msg.Chat.ID, text.String()), nil
		} else {
			id, err := strconv.ParseInt(args, 10, 0)
			if err != nil {
				return tgbotapi.MessageConfig{}, err
			}
			queryId = id
		}
	}

	b, err := queryBalance(ctx, queryId)
	if err != nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("failed to query balance: %s", err)
	}
	text.WriteString(fmt.Sprintf("当前余额：%d\n", b))

	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())
	return resp, nil
}

func queryBalance(ctx context.Context, id int64) (int64, error) {
	return db.Balance(ctx, id)
}

func (Balance) Help() string {
	return "用户的猪猪币余额\n\n用法：/balance [user_id]\n\n如果不提供user_id则查询当前用户，只有银行员工可以查询其他用户"
}
