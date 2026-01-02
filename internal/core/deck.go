package core

type Deck struct {
	CardStack
}

var cardDefinitions = []struct {
	face  CardFace
	value int
}{
	{Ace, 14}, // Ace high in this game
	{King, 13},
	{Queen, 12},
	{Jack, 11},
	{Number, 10}, {Number, 9}, {Number, 8}, {Number, 7}, {Number, 6},
	{Number, 5}, {Number, 4}, {Number, 3}, {Number, 2},
}

// NewDeck creates a standard 52-card deck
func NewDeck() *Deck {
	cards := make([]Card, 0, 52)
	suits := []CardSuit{Hearts, Diamonds, Clubs, Spades}
	for _, suit := range suits {
		for _, def := range cardDefinitions {
			cards = append(cards, Card{
				Face:  def.face,
				Suit:  suit,
				Value: def.value,
			})
		}
	}

	return &Deck{
		CardStack: CardStack{Cards: cards},
	}
}

// Draw removes and returns the top card from the deck
func (d *Deck) Draw() (Card, error) {
	return d.PopTop()
}

