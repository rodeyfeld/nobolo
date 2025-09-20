package app

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Run starts the ebiten game loop
func Run() {
	game := NewSimpleGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("NOBOLO - Simple ECS Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
