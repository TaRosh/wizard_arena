package main

import "github.com/TaRosh/wizard_arena/game"

func main() {
	g := game.NewGame(800, 600)
	err := g.Run()
	if err != nil {
		panic(err)
	}
}
