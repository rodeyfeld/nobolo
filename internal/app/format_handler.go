package app

import (
	"fmt"
	"nobolo/internal/core"
)

func formatCard(card core.Card) string {
	faceString := card.Face.String()
	suitString := card.Suit.String()
	if faceString == "" {
		faceString = fmt.Sprintf("%d", card.Value)
	}
	return fmt.Sprintf("%s of %s", faceString, suitString)
}
