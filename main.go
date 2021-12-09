package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	gol "playground/gameoflife"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := gol.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
