package main

import (
	"log"

	cmd "github.com/pmaru-top/telegetid/commands"
	tele "gopkg.in/telebot.v4"
)

var username string

func login(settings *tele.Settings) {
	b, err := tele.NewBot(*settings)
	if err != nil {
		log.Fatal(err)
		return
	}
	username = b.Me.Username
	log.Printf("%s Logged\n", username)

	registerHandlers(b)
	b.Start()
}

func registerHandlers(bot *tele.Bot) {
	global := bot.Group()

	global.Use(func(hf tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			sender := ctx.Message().Sender
			nickName := sender.FirstName + sender.LastName
			userName := sender.Username
			userId := sender.ID
			msg := ctx.Message().Text

			log.Printf(
				"msg=%s\t nickName=%s\t userName=%s\t userId=%d\t\n",
				msg, nickName, userName, userId,
			)

			return hf(ctx)
		}
	})

	global.Handle("/start", cmd.OnStart)
}
