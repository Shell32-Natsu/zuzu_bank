package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Help struct{}

func (h Help) Run(msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := strings.Builder{}
	for c, h := range commandMap {
		text.WriteString(fmt.Sprintf("/%s - %s\n", c, h.Help()))
	}
	resp := tgbotapi.NewMessage(msg.Chat.ID, text.String())

	return resp, nil
}

func (h Help) Help() string {
	return "Show the helps for all commands"
}
