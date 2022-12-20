package main

import (
	"Max-Inventary/settings"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		),

		fx.Invoke(
			func(s *settings.Settings) {
				log.Print(s)
			},
		),
	)
	app.Run()
}
