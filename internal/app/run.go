package app

import (
	"image/color"
	"log"
	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

// Update implements ebiten.Game
func (gh *GameHandler) Update() error {
	if gh.GameState == core.GameStateGameRunning {
		// Playing a card
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gh.Game.handleCardThrow()
		}

		// Slap attempt
		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			gh.Game.handleSlap()
		}

		// Win check every tick
		_ = gh.Game.isGameWon()
	}

	// Simple click to start/restart
	if (gh.GameState == core.UnknownGameState || gh.GameState == core.GameStateGameOver) &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gh.GameState = core.GameStateGameRunning
		gh.Game.startGame()
	}

	return nil
}

// Draw implements ebiten.Game
func (gh *GameHandler) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{24, 28, 36, 255})
	drawGame(screen, gh.Game)
}

// Layout implements ebiten.Game
func (gh *GameHandler) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
