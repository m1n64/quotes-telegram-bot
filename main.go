package main

import (
	"cytati/services"
	"cytati/system"
	"time"
)

func main() {
	system.LoadENV()

	quoteService := services.QuoteService{}

	var quotes []services.Quote

	index := 0

	for {
		quote, err := quoteService.GetRandomQuote()
		if err == nil || quote.QuoteAuthor == "" {
			quotes = append(quotes, quote)
			index++
		}

		if index == 2 {
			break
		}

		time.Sleep(time.Second * 3)
	}

	telegramBotService := services.TelegramService{}

	telegramBotService.InitBot()
	telegramBotService.SendQuotes(quotes)
}
