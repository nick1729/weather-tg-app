package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = fmt.Sprint("Available commands:",
					"\n/weather - show current weather",
					"\n/city - set name of the city",
					"\n/coordinates - set coordinates",
					"\n/units - set measurement units",
					"\n/lang - set language")
			case "weather":
				msg.Text = "Loading weather data..."
			case "city":
				c.City = update.Message.CommandArguments()
			case "coordinates":
				s := update.Message.CommandArguments()
				args := strings.Split(s, ", ")
				if len(args) == 2 {
					c.Coord.Lon, _ = strconv.ParseFloat(args[0], 64)
					c.Coord.Lat, _ = strconv.ParseFloat(args[1], 64)
				}
				msg.Text = fmt.Sprintf("Lon %f, Lat %f", c.Coord.Lon, c.Coord.Lat)
			case "units":
				c.Units = update.Message.CommandArguments()
				msg.Text = fmt.Sprintf("Units = %s", c.Units)
			case "lang":
				c.Lang = update.Message.CommandArguments()
				msg.Text = fmt.Sprintf("Lang = %s", c.Lang)
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
