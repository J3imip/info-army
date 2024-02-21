package main

import (
	"net/http"

	"github.com/J3imip/info-army/config"
	"github.com/J3imip/info-army/handlers"
	"github.com/J3imip/info-army/telegram_client"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"
)

const basePath = "/api/v1"

type App struct {
	log    *logan.Entry
	config config.Config
}

func (a *App) Run() {
	a.log.Info("Starting request processing")

	r := chi.NewRouter()
	r.Use(
		ape.LoganMiddleware(a.log),
		ape.CtxMiddleware(
			handlers.CtxLog(a.log),
			handlers.CtxTelegramClient(telegram_client.New(a.config)),
			handlers.CtxBotConfig(a.config.BotConfig()),
		),
	)

	r.Route(basePath, func(r chi.Router) {
		r.Post("/update", handlers.Update)
	})

	if err := http.ListenAndServe(a.config.BotConfig().Port, r); err != nil {
		a.log.WithError(err).Error("Failed to start server")
	}
}
