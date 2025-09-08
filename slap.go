package main

// SlapType represents the different types of valid slaps
type SlapType int

const (
	NoSlap       SlapType = iota
	ThreeInOrder          // Three cards in numeric order
	QueenKing             // Queen followed by King
	AddToTen              // Two cards adding up to ten (numbered cards only)
	Sandwich              // Two identical cards separated by one card
	Doubles               // Two identical cards in order
)

// String returns the string representation of SlapType
func (st SlapType) String() string {
	switch st {
	case ThreeInOrder:
		return "Three in Order"
	case QueenKing:
		return "Queen-King"
	case AddToTen:
		return "Add to Ten"
	case Sandwich:
		return "Sandwich"
	case Doubles:
		return "Doubles"
	default:
		return "No Slap"
	}
}

// CheckForSlap examines the top cards of the pile to see if a slap is valid
// Returns the type of slap found, or NoSlap if none
// Priority order: Doubles > Queen-King > Sandwich > Three-in-Order > Add-to-Ten
func CheckForSlap(pile []Card) SlapType {
	if len(pile) < 2 {
		return NoSlap
	}

	// Check doubles first (highest priority)
	if checkDoubles(pile) {
		return Doubles
	}

	// Check Queen-King
	if checkQueenKing(pile) {
		return QueenKing
	}

	// Check sandwich (need at least 3 cards) - higher priority than add to ten
	if len(pile) >= 3 && checkSandwich(pile) {
		return Sandwich
	}

	// Check three in order (need at least 3 cards)
	if len(pile) >= 3 && checkThreeInOrder(pile) {
		return ThreeInOrder
	}

	// Check add to ten last (lowest priority, only for exactly 2 cards)
	if len(pile) == 2 && checkAddToTen(pile) {
		return AddToTen
	}

	return NoSlap
}

// checkDoubles checks if the top two cards are identical
func checkDoubles(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}

	top := pile[len(pile)-1]
	second := pile[len(pile)-2]

	return top.Face == second.Face && top.Value == second.Value
}

// checkQueenKing checks if the top card is King and second is Queen
func checkQueenKing(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}

	top := pile[len(pile)-1]
	second := pile[len(pile)-2]

	return second.Face == Queen && top.Face == King
}

// checkAddToTen checks if two numbered cards add up to 10
func checkAddToTen(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}

	top := pile[len(pile)-1]
	second := pile[len(pile)-2]

	// Only numbered cards count
	if top.Face != Number || second.Face != Number {
		return false
	}

	return top.Value+second.Value == 10
}

// checkSandwich checks if two identical cards are separated by one card
func checkSandwich(pile []Card) bool {
	if len(pile) < 3 {
		return false
	}

	top := pile[len(pile)-1]
	third := pile[len(pile)-3]

	return top.Face == third.Face && top.Value == third.Value
}

// checkThreeInOrder checks if three cards are in numeric sequence
func checkThreeInOrder(pile []Card) bool {
	if len(pile) < 3 {
		return false
	}

	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	third := pile[len(pile)-3]

	// Get values for comparison
	values := []int{third.Value, second.Value, top.Value}

	// Check if they're in ascending order
	if values[0]+1 == values[1] && values[1]+1 == values[2] {
		return true
	}

	// Check if they're in descending order
	if values[0]-1 == values[1] && values[1]-1 == values[2] {
		return true
	}

	return false
}
