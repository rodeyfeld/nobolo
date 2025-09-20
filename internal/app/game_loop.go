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

// GameLoop manages turn progression, challenges, and slaps.
// Simple input:
// - Space: next player plays a card
// - S: slap (current player attempts a slap)
func (g *SimpleGame) GameLoop() {
	// Initialize if needed
	if g.GameState == core.UnknownGameState {
		g.startGame()
	}

	var pile = &g.Pile

	// basic tick - called from Update using input triggers, not auto-run
	// Playing a card
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.GameState == core.GameStateGameRunning {
		// if someone is out of cards, skip them
		if g.Players[g.CurrentPlayer].HandSize() == 0 {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			return
		}

		card, err := g.Players[g.CurrentPlayer].PlayTopCard()
		if err != nil {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			return
		}
		pile.Cards = append(pile.Cards, card)
		g.appendLog(fmt.Sprintf("%s played %d of %d", g.Players[g.CurrentPlayer].Name, card.Value, card.Suit))

		// Check if card is face card to start/continue a challenge
		if chances, ok := faceChances[card.Face]; ok {
			g.challengeOwner = g.CurrentPlayer
			g.remainingChances = chances
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			return
		}

		// If in challenge, decrement chances for next player
		if g.challengeOwner != -1 {
			g.remainingChances--
			if g.remainingChances <= 0 {
				// Challenge failed: pile goes to challengeOwner
				cards := make([]core.Card, len(pile.Cards))
				copy(cards, pile.Cards)
				pile.Cards = pile.Cards[:0]
				g.Players[g.challengeOwner].AddCardsToBottom(cards)
				g.appendLog(fmt.Sprintf("Challenge: %s takes %d cards", g.Players[g.challengeOwner].Name, len(cards)))
				owner := g.challengeOwner
				g.challengeOwner = -1
				g.remainingChances = 0
				// Next player is after the challenge owner
				g.CurrentPlayer = (owner + 1) % len(g.Players)
				return
			}
		}

		// normal progression
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		return
	}

	// Slap attempt
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && g.GameState == core.GameStateGameRunning {
		slapType := core.CheckForSlap(pile.Cards)
		if slapType != core.NoSlap {
			// Success: current player gets pile
			cards := make([]core.Card, len(pile.Cards))
			copy(cards, pile.Cards)
			pile.Cards = pile.Cards[:0]
			g.Players[g.CurrentPlayer].AddCardsToBottom(cards)
			g.appendLog(fmt.Sprintf("%s slapped: %s (+%d)", g.Players[g.CurrentPlayer].Name, slapType, len(cards)))
			// reset challenge
			g.challengeOwner = -1
			g.remainingChances = 0
		} else {
			// Penalty: remove two cards from player to pile bottom
			penaltyCards, _ := g.Players[g.CurrentPlayer].RemoveTopCards(2)
			pile.Cards = append(pile.Cards, penaltyCards...)
			g.appendLog(fmt.Sprintf("%s bad slap (-%d)", g.Players[g.CurrentPlayer].Name, len(penaltyCards)))
		}
		return
	}

	// Win check every tick
	_ = g.checkWinCondition()

	// small sleep to avoid spamming logs if keys are held on some systems
	time.Sleep(5 * time.Millisecond)
}
