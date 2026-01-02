package app

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Run starts the ebiten game loop
func Run() {
	gh := NewGameHandler()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("NOBOLO")
	if err := ebiten.RunGame(gh); err != nil {
		log.Fatal(err)
	}
}
