package main

import (
	"log"

	"github.com/matheusrb95/20-games-challenge/pong/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.ScreenWidth*2, constants.ScreenHeight*2)
	ebiten.SetWindowTitle("Pong")

	g := NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
