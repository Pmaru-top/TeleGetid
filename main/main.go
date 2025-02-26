package main

import (
	"log"
	"time"

	"github.com/pmaru-top/telegetid/config"
	_ "github.com/pmaru-top/telegetid/logger" //init logger
	"github.com/pmaru-top/telegetid/util"

	tele "gopkg.in/telebot.v4"
)

const CONFIG_PATH = "./config.json"

func main() {
	config, err := config.ReadOrCreateConfig(CONFIG_PATH)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	util.SetupProxy(config.Proxy)

	settings := &tele.Settings{
		Token:   config.Token,
		Updates: 100,
		Poller: &util.AutoReconnectPoller{
			BasePoller: &tele.LongPoller{
				Timeout: 5 * time.Second,
			},
		},
		OnError: func(err error, context tele.Context) {
			log.Fatalln(err)
		},
	}

	login(settings)
}
