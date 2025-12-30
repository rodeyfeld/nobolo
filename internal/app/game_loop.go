package app

import (
	"fmt"
	"time"

	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Challenge counters for face cards
var faceChances = map[core.CardFace]int{
	core.Jack:  1,
	core.Queen: 2,
	core.King:  3,
	core.Ace:   4,
}



func (g *SimpleGame) GameLoop() {
	if g.GameState == core.UnknownGameState {
		g.startGame()
	}

	// basic tick - called from Update using input triggers, not auto-run
	// Playing a card
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.GameState == core.GameStateGameRunning {
		
		currentPlayer := &g.Players[g.CurrentPlayer]
		
		if len(currentPlayer.Hand) == 0 {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			return
		}

		card, err := g.Players[g.CurrentPlayer].PlayTopCard()
		if err != nil {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			return
		}
		g.Pile = append(g.Pile, card)
		g.appendLog(fmt.Sprintf("%s played %s", g.Players[g.CurrentPlayer].Name, formatCard(card)))

		// Check if card is face card to start/continue a challenge
		if _, ok := faceChances[card.Face]; ok {
			g.handleFaceCard(card)
			return
		}

		// If in challenge, decrement chances for current player
		if g.progressChallenge() {
			return
		}

		// normal progression (not in challenge)
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		return
	}

	// Slap attempt
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && g.GameState == core.GameStateGameRunning {
		g.handleSlap()
		return
	}

	// Win check every tick
	_ = g.checkWinCondition()

	// small sleep to avoid spamming logs if keys are held on some systems
	time.Sleep(5 * time.Millisecond)
}
