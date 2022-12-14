package bot

import (
	"context"
	"fmt"
	"unicode/utf16"

	"github.com/Shell32-Natsu/zuzu_bank/internal/commands"
	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func utf16Length(s string) int {
	return len(utf16.Encode([]rune(s)))
}

func ParseMessage(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.Chattable, error) {
	if !config.IsAllowedUser(msg.From.ID) {
		return tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("User id %d is not allowed.", msg.From.ID)), nil
	}

	command := msg.Command()
	if command == "" {
		return nil, nil
	}

	resp, err := commands.RunCommand(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to run command: %s", err)
	}
	resp.ReplyToMessageID = msg.MessageID
	resp.Entities = append(resp.Entities, tgbotapi.MessageEntity{
		Type:   "code",
		Offset: 0,
		Length: utf16Length(resp.Text),
	})

	return resp, nil
}
