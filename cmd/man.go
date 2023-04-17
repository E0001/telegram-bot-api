package main

import (
	"fmt"
	tgbotapi "github.com/E0001/telegram-bot-api/v5"
	"log"
)

// 5730487425:AAH3LSGVDGC7qcgYwoJP5dn_8C-xLGMwWik  -1001604760179

/*

https://api.telegram.org/bot5730487425:AAH3LSGVDGC7qcgYwoJP5dn_8C-xLGMwWik/sendMessage?chat_id=@<super_group_name>/<topic_name>&text=<message_text>
https://api.telegram.org/bot5730487425:AAH3LSGVDGC7qcgYwoJP5dn_8C-xLGMwWik/sendMessage?chat_id=-1001604760179/#test&text="你好"


*/

func main() {
	bot, err := tgbotapi.NewBotAPI("5730487425:AAH3LSGVDGC7qcgYwoJP5dn_8C-xLGMwWik")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	msg := tgbotapi.NewMessage(-1001604760179, "#test 你好 **************************")

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("test", "https://example.com"),
		),
	)

	msg.MessageThreadId = 41186

	bot.Send(msg)

	return

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		fmt.Println(update)

		if !update.Message.IsCommand() { // ignore any non-command Messages
			// continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
