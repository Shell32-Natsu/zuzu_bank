package bot

import (
	"fmt"

	"github.com/Shell32-Natsu/zuzu_bank/internal/commands"
	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ParseMessage(c *config.Config, msg *tgbotapi.Message) (tgbotapi.Chattable, error) {
	if !c.IsAllowedUser(msg.From.ID) {
		return tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("User id %d is not allowed.", msg.From.ID)), nil
	}

	command := msg.Command()
	if command == "" {
		return nil, nil
	}

	resp, err := commands.RunCommand(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to run command: %s", err)
	}
	resp.ReplyToMessageID = msg.MessageID

	return resp, nil
}
