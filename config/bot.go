package config

import (
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Bot interface {
	BotConfig() BotConfig
}

type BotConfig struct {
	ApiToken string   `fig:"api_token,required"`
	ApiURL   *url.URL `fig:"api_url,required"`
	Port     string   `fig:"port,required"`
	Channel  string   `fig:"channel,required"`
	Photo    *url.URL `fig:"photo,required"`
}

type bot struct {
	once   comfig.Once
	getter kv.Getter
}

func NewBot(getter kv.Getter) Bot {
	return &bot{
		getter: getter,
	}
}

func (c *bot) BotConfig() BotConfig {
	return c.once.Do(func() interface{} {
		var cfg BotConfig

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "bot")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out bot"))
		}

		return cfg
	}).(BotConfig)
}
