package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

var charsToEscape = [...]string{"*", "`", "_", "[", "]"}

func EscapeMarkdown(text string) string {
	for _, char := range charsToEscape {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}

func GetTextForNewMessage(update tgbotapi.Update) string {
	var text = ""
	if update.Message.ForwardFromChat != nil {
		var chat = update.Message.ForwardFromChat
		text = GetTextForChat(chat)
	} else {
		var user = &tgbotapi.User{}
		if update.Message.ForwardFrom != nil {
			user = update.Message.ForwardFrom
		} else {
			user = update.Message.From
		}
		text = GetTextForUser(user)
	}
	return text
}

func GetTextForUser(user *tgbotapi.User) string {
	var sb strings.Builder
	var username = user.UserName
	if username == "" {
		username = "No username"
	} else {
		username = "@" + EscapeMarkdown(username)
	}
	var lastName = user.LastName
	var firstName = user.FirstName

	if lastName == "" {
		lastName = "No Last Name"
	} else {
		lastName = EscapeMarkdown(lastName)
	}

	if firstName == "" {
		firstName = "No First Name"
	} else {
		firstName = EscapeMarkdown(firstName)
	}
	sb.WriteString("*User info:*\n")
	sb.WriteString("*ID:* `" + strconv.FormatInt(int64(user.ID), 10) + "`\n")
	sb.WriteString("*Username:* " + username + "\n")
	sb.WriteString("*First Name:* " + firstName + "\n")
	sb.WriteString("*Last Name:* " + lastName)
	return sb.String()
}

func GetTextForChat(chat *tgbotapi.Chat) string {
	var sb strings.Builder
	var username = chat.UserName
	if username == "" {
		username = "No username"
	} else {
		username = "@" + EscapeMarkdown(username)
	}
	var title = chat.Title

	if title == "" {
		title = "No title"
	} else {
		title = EscapeMarkdown(title)
	}
	chatType := chat.Type
	chatType = strings.ToUpper(string(chatType[0])) + string(chatType[1:])
	sb.WriteString("*" + chatType + " info:*\n")
	sb.WriteString("*ID:* `" + strconv.FormatInt(chat.ID, 10) + "`\n")
	sb.WriteString("*Username:* " + username + "\n")
	sb.WriteString("*Title:* " + title)
	return sb.String()
}
