package main

import (
	"log"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var c tConfig

func init() {

	var errCfg error

	c, errCfg = loadCfg("./config/config.json")
	if errCfg != nil {
		log.Fatal(errCfg)
	}
}

func main() {

	bot, errBot := tgbotapi.NewBotAPI(c.Token)
	if errBot != nil {
		log.Fatal(errBot)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, errUpd := bot.GetUpdatesChan(u)
	if errUpd != nil {
		log.Print(errUpd)
	}

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
