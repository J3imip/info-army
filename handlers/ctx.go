package handlers

import (
	"context"
	"net/http"

	"github.com/J3imip/info-army/config"
	"github.com/J3imip/info-army/telegram_client"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	telegramClienterCtxKey
	botConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxTelegramClient(telegramClient *telegram_client.TelegramClienter) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, telegramClienterCtxKey, telegramClient)
	}
}

func TelegramClient(r *http.Request) *telegram_client.TelegramClienter {
	return r.Context().Value(telegramClienterCtxKey).(*telegram_client.TelegramClienter)
}

func CtxBotConfig(botConfig config.BotConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, botConfigCtxKey, botConfig)
	}
}

func BotConfig(r *http.Request) config.BotConfig {
	return r.Context().Value(botConfigCtxKey).(config.BotConfig)
}
