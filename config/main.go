package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	Bot
	comfig.Logger
}

type config struct {
	getter kv.Getter
	comfig.Logger
	Bot
}

func New(getter kv.Getter) Config {
	return &config{
		getter: getter,
		Bot:    NewBot(getter),
		Logger: comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
