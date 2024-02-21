package telegram_client

import (
	bytes2 "bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/J3imip/info-army/config"
	"github.com/J3imip/info-army/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	sendMessageURL = "/sendMessage"
	sendPhotoURl   = "/sendPhoto"
)

type TelegramClienter struct {
	ApiToken string
	ApiURL   *url.URL
	log      *logan.Entry
}

func New(cfg config.Config) *TelegramClienter {
	return &TelegramClienter{
		ApiToken: cfg.BotConfig().ApiToken,
		ApiURL:   cfg.BotConfig().ApiURL,
		log:      cfg.Log(),
	}
}

func (t *TelegramClienter) SendMessage(
	chatId int64,
	text string,
	parseMode *string,
	replyMarkup json.Marshaler,
	entities ...types.Entities,
) error {
	if len(text) == 0 {
		return errors.New("text is empty")
	}

	request := types.SendMessageRequest{
		ChatID: chatId,
		Text:   text,
	}

	var err error
	if replyMarkup != nil {
		request.ReplyMarkup, err = json.Marshal(replyMarkup)
		if err != nil {
			return errors.Wrap(err, "failed to marshal reply markup")
		}
	}
	if parseMode != nil {
		request.ParseMode = *parseMode
	}
	if len(entities) > 0 {
		request.Entities = entities
	}

	var buf bytes2.Buffer
	err = json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return errors.Wrap(err, "failed to encode request")
	}

	resp, err := http.Post(t.ApiURL.String()+t.ApiToken+sendMessageURL, "application/json", &buf)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	if resp.StatusCode != http.StatusOK {
		var tgErr types.TelegramError
		if err := json.NewDecoder(resp.Body).Decode(&tgErr); err != nil {
			return errors.Wrap(err, "failed to decode error")
		}

		return errors.Wrap(errors.New(tgErr.Description), "telegram error")
	}

	return nil
}

func (t *TelegramClienter) SendPhoto(
	chatId int64,
	photo string,
	caption *string,
	parseMode *string,
	replyMarkup json.Marshaler,
	entities ...types.Entities,
) error {
	if len(photo) == 0 {
		return errors.New("photo is empty")
	}

	request := types.SendPhotoRequest{
		ChatID: chatId,
		Photo:  photo,
	}

	var err error
	if replyMarkup != nil {
		request.ReplyMarkup, err = json.Marshal(replyMarkup)
		if err != nil {
			return errors.Wrap(err, "failed to marshal reply markup")
		}
	}
	if parseMode != nil {
		request.ParseMode = *parseMode
	}
	if len(entities) > 0 {
		request.Entities = entities
	}
	if caption != nil {
		request.Caption = *caption
	}

	var buf bytes2.Buffer
	err = json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return errors.Wrap(err, "failed to encode request")
	}

	resp, err := http.Post(t.ApiURL.String()+t.ApiToken+sendPhotoURl, "application/json", &buf)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	if resp.StatusCode != http.StatusOK {
		var tgErr types.TelegramError
		if err := json.NewDecoder(resp.Body).Decode(&tgErr); err != nil {
			return errors.Wrap(err, "failed to decode error")
		}

		return errors.Wrap(errors.New(tgErr.Description), "telegram error")
	}

	return nil
}
