package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var keyboard = tgbotapi.NewKeyboardButton("start")

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file:", err)
	}
}

func main() {
	// Bot initialization
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal("error running Telegram API: ", err)
	}
	// bot.Debug = true

	fmt.Println("Authorized on account", bot.Self.UserName)

	// Bot configuration
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			fmt.Println("no messages")
			continue
		}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = keyboard

	if update.Message.IsCommand() {
		msg.Text = handleCommands(update.Message.Command())
	} else {
		msg.Text = fmt.Sprintf("Hello, %s!", update.Message.From.FirstName)
	}

	if _, err := bot.Send(msg); err != nil {
		log.Fatalln(err)
	}

	}

}

func handleCommands(command string) string {
	switch command {
	case "members":
		return "Getting members"
	case "teams":
		return "Getting teams"
	case "creators":
		return "Getting creators"
	default:
		return "I don't know this command"
	}
}