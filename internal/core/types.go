package core

import "errors"

// Basic game types

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

type Pile struct {
	Cards []Card
}

type Player struct {
	Name string
	Hand []Card
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: make([]Card, 0),
	}
}

func (p *Player) PlayTopCard() (Card, error) {
	if len(p.Hand) == 0 {
		return Card{}, errors.New("player has no cards")
	}
	idx := len(p.Hand) - 1
	c := p.Hand[idx]
	p.Hand = p.Hand[:idx]
	return c, nil
}

func (p *Player) AddCardsToBottom(cards []Card) {
	if len(cards) == 0 {
		return
	}
	p.Hand = append(cards, p.Hand...)
}

// RemoveTopCards removes n cards from the top (end) of the hand
func (p *Player) RemoveTopCards(n int) ([]Card, error) {
	if n <= 0 {
		return []Card{}, nil
	}
	if len(p.Hand) == 0 {
		return []Card{}, nil
	}
	if len(p.Hand) < n {
		cards := make([]Card, len(p.Hand))
		copy(cards, p.Hand)
		p.Hand = p.Hand[:0]
		return cards, nil
	}
	start := len(p.Hand) - n
	cards := make([]Card, n)
	copy(cards, p.Hand[start:])
	p.Hand = p.Hand[:start]
	return cards, nil
}

type GameState byte

const (
	UnknownGameState = GameState(iota)
	GameStateGameRunning
	GameStateGameOver
)

// String returns the string representation of GameState
func (gs GameState) String() string {
	switch gs {
	case GameStateGameRunning:
		return "Running"
	case GameStateGameOver:
		return "Game Over"
	default:
		return "Unknown"
	}
}
