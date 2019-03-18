package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type config struct {
	botToken string
}

func getConfig() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. You need to create .env file and fill BOT_TOKEN param")
	}
	configInstance := config{os.Getenv("BOT_TOKEN")}
	return configInstance
}

func main() {
	config := getConfig()
	bot, err := tgbotapi.NewBotAPI(config.botToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		messageText := update.Message.Text
		if messageText == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi! I can show you some info about users, chats or channels.\nJust forward message to me or send username like @Durov :)")
			bot.Send(msg)
		} else if messageText != "" && string(messageText[0]) == "@" && strings.Index(messageText, " ") == -1 {
			chat, err := bot.GetChat(tgbotapi.ChatConfig{SuperGroupUsername: messageText})
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Sorry, i can't find chat with username "+messageText)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					GetTextForChat(&chat))
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "markdown"
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTextForNewMessage(update))
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "markdown"
			bot.Send(msg)
		}
	}
}
