package main

import (
	"shoot/game"

	"github.com/firefly-zero/firefly-go/firefly"
)

func init() {
	firefly.Boot = game.Boot
	firefly.Update = game.Update
	firefly.Render = game.Render
}
