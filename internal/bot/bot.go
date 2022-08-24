package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/Shell32-Natsu/zuzu_bank/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(ctx context.Context) error {
	bot, err := tgbotapi.NewBotAPI(config.BotKey())
	if err != nil {
		return fmt.Errorf("failed to get new bot: %s", err)
	}

	bot.Debug = config.IsDebug()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			resp, err := ParseMessage(ctx, update.Message)
			if err != nil {
				if config.IsAdmin(update.Message.From.ID) {
					r := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					r.ReplyToMessageID = update.Message.MessageID
					_, newErr := bot.Send(r)
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
