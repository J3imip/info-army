package handlers

import (
	"fmt"
	"net/http"

	"github.com/J3imip/info-army/types"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Start(request *types.Update, w http.ResponseWriter, r *http.Request) {
	if err := TelegramClient(r).SendMessage(
		request.Message.Chat.ID,
		fmt.Sprintf("Привіт, %s 👋! Тебе вітає команда INFOAрмії, долучайся до нас. "+
			"Але спочатку декілька простих питань.", request.Message.From.FirstName),
		nil,
		nil,
	); err != nil {
		Log(r).WithError(err).Error("failed to send message")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err := TelegramClient(r).SendMessage(
		request.Message.Chat.ID,
		"Готовий проводити спецоперації в російських соцмережах та впливати на перебіг війни?",
		nil,
		&types.ReplyMarkup{
			InlineKeyboardMarkup: [][]types.InlineKeyboardButton{
				{
					{
						Text:         "Так✅",
						CallbackData: "ready",
					},
					{
						Text:         "Ні❌",
						CallbackData: "not_ready",
					},
				},
			},
		},
	); err != nil {
		Log(r).WithError(err).Error("failed to send message")
		ape.RenderErr(w, problems.InternalError())
		return
	}
}
