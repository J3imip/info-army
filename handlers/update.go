package handlers

import (
	"net/http"

	"github.com/J3imip/info-army/requests"
	"github.com/J3imip/info-army/types"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Update(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdate(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	switch {
	case request.Message != nil:
		if err := handleCommands(request, w, r); err != nil {
			Log(r).WithError(err).Error("failed to handle command")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	case request.CallbackQuery != nil:
		if err := handleCallbackQuery(request, w, r); err != nil {
			Log(r).WithError(err).Error("failed to handle callback query")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

}

func handleCommands(request *types.Update, w http.ResponseWriter, r *http.Request) error {
	if len(request.Message.Entities) > 0 &&
		request.Message.Entities[0].Type != "bot_command" {
		return nil
	}

	switch request.Message.Text {
	case "/start":
		Start(request, w, r)
	default:
		if err := TelegramClient(r).SendMessage(
			request.Message.Chat.ID,
			"Я не знаю такої команди 😔",
			nil,
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}

func handleCallbackQuery(request *types.Update, w http.ResponseWriter, r *http.Request) error {
	switch request.CallbackQuery.Data {
	case "not_ready":
		if err := TelegramClient(r).SendMessage(
			request.CallbackQuery.Message.Chat.ID,
			"Русні пизда",
			nil,
			nil,
		); err != nil {
			return err
		}
	case "ready":
		if err := TelegramClient(r).SendMessage(
			request.CallbackQuery.Message.Chat.ID,
			"Хочеш аби в рф більше палало? 🔥🔥",
			nil,
			&types.ReplyMarkup{
				InlineKeyboardMarkup: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "Так✅",
							CallbackData: "yes",
						},
						{
							Text:         "Ні❌",
							CallbackData: "no",
						},
					},
				},
			},
		); err != nil {
			return err
		}
	case "yes":
		if err := TelegramClient(r).SendMessage(
			request.CallbackQuery.Message.Chat.ID,
			"Супер! Долучайся до нас та стань частиною INFOАрмії \U0001FAE1\n"+
				"Переходь в наш основний канал, там буде багато вогняних подій - "+BotConfig(r).Channel,
			nil,
			nil,
		); err != nil {
			return err
		}
	case "no":
		if err := TelegramClient(r).SendPhoto(
			request.CallbackQuery.Message.Chat.ID,
			BotConfig(r).Photo.String(),
			nil,
			nil,
			nil,
		); err != nil {
			return err
		}
	}

	return nil
}
