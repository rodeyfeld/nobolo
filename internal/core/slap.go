package core

// SlapType represents the different types of valid slaps
type SlapType int

const (
	NoSlap       SlapType = iota
	ThreeInOrder          // Three cards in numeric order
	QueenKing             // Queen Followed by King
	KingQueen             // King Followed by Queen
	AddToTen              // Two cards adding up to ten (numbered cards only)
	Sandwich              // Two identical cards separated by one card
	Doubles               // Two identical cards in order
)

func (st SlapType) String() string {
	switch st {
	case ThreeInOrder:
		return "Three in Order"
	case QueenKing:
		return "Queen-King"
	case KingQueen:
		return "King-Queen"
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

func CheckSlapType(pile *Pile) SlapType {
	cards := pile.Cards
	if len(cards) < 2 {
		return NoSlap
	}

	if isDoubles(cards) {
		return Doubles
	}
	if isQueenKing(cards) {
		return QueenKing
	}
	if isKingQueen(cards) {
		return KingQueen
	}
	if isSandwich(cards) {
		return Sandwich
	}
	if isThreeInOrder(cards) {
		return ThreeInOrder
	}
	if isAddToTen(cards) {
		return AddToTen
	}
	return NoSlap
}

func isDoubles(cards []Card) bool {
	if len(cards) < 2 {
		return false
	}
	top := cards[len(cards)-1]
	second := cards[len(cards)-2]
	return top.Face == second.Face && top.Value == second.Value
}
func isKingQueen(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	return top.Face == Queen && second.Face == King
}

func isQueenKing(pile []Card) bool {
	if len(pile) < 2 {
		return false
	}
	top := pile[len(pile)-1]
	second := pile[len(pile)-2]
	return second.Face == Queen && top.Face == King
}

func isAddToTen(pile []Card) bool {
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

func isSandwich(pile []Card) bool {
	if len(pile) < 3 {
		return false
	}
	top := pile[len(pile)-1]
	third := pile[len(pile)-3]
	return top.Face == third.Face && top.Value == third.Value
}

func isThreeInOrder(pile []Card) bool {
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
