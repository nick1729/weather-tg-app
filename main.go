package main

import (
	"log"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	c      tConfig
	w      tWeather
	err    error
	byCity bool
)

func init() {

	var errCfg error

	c, errCfg = loadCfg("./config/config.json")
	if errCfg != nil {
		log.Fatal(errCfg)
	}

	byCity = true
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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = printHelp()
			case "weather":
				w, err = getWeather(c, byCity)
				msg.Text = printWeather(c, w, err)
			case "city":
				c.City = update.Message.CommandArguments()
				byCity = true
			case "coordinates":
				c, byCity, msg.Text = setCoordinates(c, update.Message.CommandArguments())
			case "units":
				c, w, msg.Text = setConvUnits(c, w, update.Message.CommandArguments())
			case "lang":
				c.Lang = update.Message.CommandArguments()
			case "exit":
				log.Println("Exit")
				os.Exit(1)
			default:
				msg.Text = "Wrong command. Type /help"
			}
		} else {
			msg.Text = "Type /help"
		}
		bot.Send(msg)
	}
}
