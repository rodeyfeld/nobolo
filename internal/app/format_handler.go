package app

import (
	"fmt"
	"nobolo/internal/core"
)

func (g *Game) appendLog(s string) {
	g.logLines = append(g.logLines, s)
	fmt.Println(s)
}

func formatCard(card core.Card) string {
	faceString := card.Face.String()
	suitString := card.Suit.String()
	if faceString == "" {
		faceString = fmt.Sprintf("%d", card.Value)
	}
	return fmt.Sprintf("%s of %s", faceString, suitString)
}
