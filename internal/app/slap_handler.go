package app

import (
	"fmt"
	"nobolo/internal/core"
)

func (g *Game) handleSlap() {
	slapType := core.CheckSlapType(g.Pile)
	if slapType != core.NoSlap {
		wonCards := len(g.Pile.Cards)
		g.Players[g.CurrentPlayer].AddCardsToBottom(g.Pile.Cards)
		g.Pile.Cards = g.Pile.Cards[:0]
		g.appendLog(fmt.Sprintf("%s slapped: %s (+%d)", g.Players[g.CurrentPlayer].Name, slapType, wonCards))
		g.challengeOwner = -1
		g.remainingChances = 0
		return
	} else {
		penaltyCards, _ := g.Players[g.CurrentPlayer].RemoveTopCards(2)
		g.Pile.Cards = append(g.Pile.Cards, penaltyCards...)
		g.appendLog(fmt.Sprintf("%s bad slap (-%d)", g.Players[g.CurrentPlayer].Name, len(penaltyCards)))
	}
}
