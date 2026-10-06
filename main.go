package main

import (
	"github.com/azulgautam79/go-discord-bot/bot"
	"github.com/azulgautam79/go-discord-bot/config"
)

func main() {

	//! Config
	cfg := config.MustLoad()

	bot.Start(cfg)

	<-make(chan struct{})
}
