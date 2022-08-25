package commands

import (
	"context"
	"fmt"
	"github.com/Shell32-Natsu/zuzu_bank/internal/logging"
	"strconv"
	"strings"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunCommand(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	for c, h := range commandMap {
		if c == msg.Command() {
			args := msg.CommandArguments()
			if args != "" && !config.IsAdmin(msg.From.ID) {
				return tgbotapi.NewMessage(msg.Chat.ID, "只有银行员工可以查询其他账户"), nil
			}
			return h.Run(ctx, msg)
		}
	}
	// Unknown command
	return tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("命令不存在: %s", msg.Command())), nil
}

func getQueryUserIdFromCommandArgs(fromId int64, args string, pos int) int64 {
	args = strings.TrimSpace(args)
	logging.LogDebugf("fromId=%d, args=%q, pos=%d", fromId, args, pos)
	if args == "" {
		return fromId
	}
	result := fromId

	argsSlice := strings.Split(args, " ")
	logging.LogDebug(argsSlice)
	if len(argsSlice) > pos {
		id, err := strconv.ParseInt(argsSlice[pos], 10, 0)
		if err != nil {
			return -1
		}
		result = id
	}
	return result
}

func writeStringf(b *strings.Builder, f string, args ...any) {
	b.WriteString(fmt.Sprintf(f, args...))
}
