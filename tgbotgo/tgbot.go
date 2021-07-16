package tgbotgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	tgtypes "github.com/incrypto32/airdropbot/tgbotgo/types"
)

type TgBot struct {
	Token  string
	Url    string
	Client *http.Client
}

func (tgbot *TgBot) GetAuthorizedUrl() string {
	return tgbot.Url + "/bot" + tgbot.Token
}

func New(token string, baseUrl string) (tgbot *TgBot, err error) {
	uri, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	baseUrl = uri.String()

	if uri.Scheme != "https" {
		return nil, errors.New("url not https")
	}

	return &TgBot{Token: token, Url: baseUrl, Client: &http.Client{}}, err
}

func (bot *TgBot) SendTextMessage(chatId int, text string) (message *tgtypes.Message, err error) {
	methodUrl := bot.GetAuthorizedUrl() + "/sendMessage"
	msg := &tgtypes.OutgoingMessage{
		ChatId: chatId, Text: text,
	}

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return message, err
	}

	req, err := http.NewRequest("POST", methodUrl, bytes.NewBuffer(msgJson))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return message, err
	}

	resp, err := bot.Client.Do(req)
	if err != nil {
		return message, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return message, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("an error occurred")
	}

	message = new(tgtypes.Message)

	err = json.Unmarshal(b, &message)

	return message, err
}
