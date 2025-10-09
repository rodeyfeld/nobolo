package app

import (
	"fmt"
	"nobolo/internal/core"
)

// formatCard returns a human-readable string like "Jack of Hearts" or "7 of Clubs".
func formatCard(card core.Card) string {
	faceString := card.Face.String()
	suitString := card.Suit.String()
	if faceString == "" {
		faceString = fmt.Sprintf("%d", card.Value)
	}
	return fmt.Sprintf("%s of %s", faceString, suitString)
}
