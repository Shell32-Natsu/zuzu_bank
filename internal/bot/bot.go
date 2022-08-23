package bot

import (
	"fmt"
	"log"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(c *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(c.BotKey)
	if err != nil {
		return fmt.Errorf("failed to get new bot: %s", err)
	}

	bot.Debug = c.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			resp, err := ParseMessage(c, update.Message)
			if err != nil {
				if c.IsAdmin(update.Message.From.ID) {
					_, newErr := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
					if newErr != nil {
						log.Panicf("original error:\n%s\nnew error:\n%s\n", err, newErr)
					}
				} else {
					log.Printf("failed to parse message: %s", err)
					continue
				}
			}

			if resp == nil {
				// No response should be sent
				continue
			}

			bot.Send(resp)
		}
	}

	return fmt.Errorf("returned from bot update loop")
}
