package core

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

// CheckForSlap examines the pile to see if a slap is valid
// Priority order: Doubles > Queen-King > Sandwich > Three-in-Order > Add-to-Ten
func CheckForSlap(pile []Card) SlapType {
	if len(pile) < 2 {
		return NoSlap
	}

	if checkDoubles(pile) {
		return Doubles
	}
	if checkQueenKing(pile) {
		return QueenKing
	}
	if len(pile) >= 3 && checkSandwich(pile) {
		return Sandwich
	}
	if len(pile) >= 3 && checkThreeInOrder(pile) {
		return ThreeInOrder
	}
	if len(pile) >= 2 && checkAddToTen(pile) {
		return AddToTen
	}
	return NoSlap
}

func checkDoubles(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	return top.Face == second.Face && top.Value == second.Value
}

func checkQueenKing(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	return second.Face == Queen && top.Face == King
}

func checkAddToTen(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	if top.Face != Number || second.Face != Number {
		return false
	}
	return top.Value+second.Value == 10
}

func checkSandwich(pile []Card) bool {
	if len(pile) < 3 {
		return false
	}
	top := pile[len(pile)-1]
	third := pile[len(pile)-3]
	return top.Face == third.Face && top.Value == third.Value
}

func checkThreeInOrder(pile []Card) bool {
	if len(pile) < 3 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	third := pile[len(pile)-3]
	values := []int{third.Value, second.Value, top.Value}
	if values[0]+1 == values[1] && values[1]+1 == values[2] {
		return true
	}
	if values[0]-1 == values[1] && values[1]-1 == values[2] {
		return true
	}
	return false
}
