package main

import (
	"github.com/J3imip/info-army/config"
	"gitlab.com/distributed_lab/kit/kv"
)

func main() {
	cfg := config.New(kv.MustFromEnv())

	app := &App{
		log:    cfg.Log(),
		config: cfg,
	}

	app.Run()
}
