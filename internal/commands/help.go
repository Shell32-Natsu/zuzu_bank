package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Help struct{}

func (h Help) Run(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}

	if msg.CommandArguments() != "" {
		for c, h := range commandMap {
			if c == strings.TrimSpace(msg.CommandArguments()) {
				return tgbotapi.NewMessage(msg.Chat.ID, h.Help()), nil
			}
		}
	}

	for c, h := range commandMap {
		text.WriteString(fmt.Sprintf("/%s - %s\n", c, strings.Split(h.Help(), "\n")[0]))
	}
	return tgbotapi.NewMessage(msg.Chat.ID, text.String()), nil
}

func (h Help) Help() string {
	return "显示帮助信息"
}
