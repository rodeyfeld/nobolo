package app

import (
	"fmt"
	"nobolo/internal/core"
)

// handleSlap processes a slap attempt by the current player.
func (g *SimpleGame) handleSlap(pile *core.Pile) {
	slapType := core.CheckForSlap(pile.Cards)
	if slapType != core.NoSlap {
		cards := make([]core.Card, len(pile.Cards))
		copy(cards, pile.Cards)
		pile.Cards = pile.Cards[:0]
		g.Players[g.CurrentPlayer].AddCardsToBottom(cards)
		g.appendLog(fmt.Sprintf("%s slapped: %s (+%d)", g.Players[g.CurrentPlayer].Name, slapType, len(cards)))
		g.challengeOwner = -1
		g.remainingChances = 0
		return
	}
	penaltyCards, _ := g.Players[g.CurrentPlayer].RemoveTopCards(2)
	pile.Cards = append(pile.Cards, penaltyCards...)
	g.appendLog(fmt.Sprintf("%s bad slap (-%d)", g.Players[g.CurrentPlayer].Name, len(penaltyCards)))
}
