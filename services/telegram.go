package services

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

type TelegramService struct {
	Bot *tgbotapi.BotAPI
}

func (service *TelegramService) InitBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		panic(err)
	}

	isDebug, err := strconv.ParseBool(os.Getenv("TELEGRAM_DEBUG"))
	if err != nil {
		panic(err)
	}

	bot.Debug = isDebug

	service.Bot = bot
}

func (service *TelegramService) SendQuotes(quotes []Quote) {
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_GROUP_ID"), 10, 64)
	if err != nil {
		panic(err)
	}

	messageText := "*Цитаты недели (кто пропустил)*:\r\n"
	for _, quote := range quotes {
		messageText += fmt.Sprintf("\"_%s_\" (C) %s\r\n\r\n", quote.QuoteText, quote.QuoteAuthor)
	}

	message := tgbotapi.NewMessage(chatID, messageText)
	message.ParseMode = tgbotapi.ModeMarkdown

	_, err = service.Bot.Send(message)
	if err != nil {
		panic(err)
	}
}
