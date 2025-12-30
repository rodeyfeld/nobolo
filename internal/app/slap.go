package app

import (
	"fmt"
	"nobolo/internal/core"
)

// handleSlap processes a slap attempt by the current player.
func (g *SimpleGame) handleSlap() {
	slapType := core.CheckForSlap(g.Pile)
	if slapType != core.NoSlap {
		wonCards := len(g.Pile)
		g.Players[g.CurrentPlayer].AddCardsToBottom(g.Pile)
		g.Pile = g.Pile[:0]
		g.appendLog(fmt.Sprintf("%s slapped: %s (+%d)", g.Players[g.CurrentPlayer].Name, slapType, wonCards))
		g.challengeOwner = -1
		g.remainingChances = 0
		return
	}
	penaltyCards, _ := g.Players[g.CurrentPlayer].RemoveTopCards(2)
	g.Pile = append(g.Pile, penaltyCards...)
	g.appendLog(fmt.Sprintf("%s bad slap (-%d)", g.Players[g.CurrentPlayer].Name, len(penaltyCards)))
}
