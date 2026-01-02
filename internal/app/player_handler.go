package app

import (
	"fmt"
	"nobolo/internal/core"
)

func (g *Game) handleCardThrow() {
	currentPlayer := &g.Players[g.CurrentPlayer]

	if currentPlayer.Hand.Count() == 0 {
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		return
	}

	card, err := g.Players[g.CurrentPlayer].Hand.PopTop()
	if err != nil {
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		return
	}
	g.Pile.Cards = append(g.Pile.Cards, card)
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

func (g *Game) handleSlap() {
	slapType := core.CheckSlapType(g.Pile)
	if slapType != core.NoSlap {
		wonCards := g.Pile.Count()
		g.Players[g.CurrentPlayer].Hand.PushBottom(g.Pile.Clear())
		g.appendLog(fmt.Sprintf("%s slapped: %s (+%d)", g.Players[g.CurrentPlayer].Name, slapType, wonCards))
		g.challengeOwner = -1
		g.remainingChances = 0
		return
	} else {
		penaltyCards := g.Players[g.CurrentPlayer].Hand.PopTopN(2)
		g.Pile.PushBottom(penaltyCards)
		g.appendLog(fmt.Sprintf("%s bad slap (-%d)", g.Players[g.CurrentPlayer].Name, len(penaltyCards)))
	}
}
