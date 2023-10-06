package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

type QuoteService struct {
}

func (service *QuoteService) GetRandomQuote() (Quote, error) {
	url := "https://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=ru"
	resp, err := http.Get(url)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Quote{}, err
		}

		var quote Quote
		err = json.Unmarshal(bodyBytes, &quote)

		if err != nil {
			return Quote{}, err
		}

		return quote, nil
	} else {
		return Quote{}, err
	}
}
