package core

// CardFace represents the face value of a card
type CardFace byte

// CardSuit represents the suit of a card
type CardSuit byte

const (
	UnknownFace = CardFace(iota)
	Jack
	Queen
	King
	Ace
	Number
)

const (
	UnknownSuit = CardSuit(iota)
	Hearts
	Diamonds
	Clubs
	Spades
)

func (cf CardFace) String() string {
	switch cf {
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return ""
	}
}

func (cs CardSuit) String() string {
	switch cs {
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	default:
		return "Unknown"
	}
}

type Card struct {
	Face  CardFace
	Suit  CardSuit
	Value int
}

