package app

import (
	"fmt"
	"nobolo/internal/core"
)

// handleFaceCard starts a challenge for the given face card.
func (g *SimpleGame) handleFaceCard(card core.Card) {
	if chances, ok := faceChances[card.Face]; ok {
		g.appendLog(fmt.Sprintf("Challenge: %s started a challenge", g.Players[g.CurrentPlayer].Name))
		g.challengeOwner = g.CurrentPlayer
		g.remainingChances = chances
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		g.appendLog(fmt.Sprintf("Challenge: %s %d chances", g.Players[g.CurrentPlayer].Name, g.remainingChances))
	}
}

// progressChallenge decrements chances for the current player and resolves if failed.
// Returns true if the turn should end (either challenge resolved or continue same player),
// and false if normal progression should occur.
func (g *SimpleGame) progressChallenge() bool {
	if g.challengeOwner == -1 {
		return false
	}
	g.remainingChances--
	g.appendLog(fmt.Sprintf("Challenge: %s has %d chances left", g.Players[g.CurrentPlayer].Name, g.remainingChances))
	if g.remainingChances > 0 {
		// same player keeps playing
		return true
	}

	// Failed challenge: give pile to owner
	g.appendLog(fmt.Sprintf("Challenge: %s failed the challenge", g.Players[g.CurrentPlayer].Name))
	cards := make([]core.Card, len(g.Pile))
	copy(cards, g.Pile)
	g.Pile = g.Pile[:0]
	g.Players[g.challengeOwner].AddCardsToBottom(cards)
	g.appendLog(fmt.Sprintf("Challenge: %s takes %d cards", g.Players[g.challengeOwner].Name, len(cards)))
	owner := g.challengeOwner
	g.challengeOwner = -1
	g.remainingChances = 0
	g.CurrentPlayer = (owner + 1) % len(g.Players)
	return true
}
