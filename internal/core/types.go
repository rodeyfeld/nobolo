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

// Card represents a single playing card
type Card struct {
	Face  CardFace
	Suit  CardSuit
	Value int
}

// Pile represents the center pile of cards
type Pile struct {
	Cards []Card
}

// Size returns the number of cards in the pile
func (p *Pile) Size() int {
	return len(p.Cards)
}

// Player represents a game player
type Player struct {
	Name string
	Hand []Card
}

// NewPlayer creates a new player with the given name
func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: make([]Card, 0),
	}
}

// AddCard adds a card to the player's hand
func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// HandSize returns the number of cards in the player's hand
func (p *Player) HandSize() int {
	return len(p.Hand)
}

// PlayTopCard removes and returns the last card (top) from the player's hand
func (p *Player) PlayTopCard() (Card, error) {
	if len(p.Hand) == 0 {
		return Card{}, errors.New("player has no cards")
	}
	idx := len(p.Hand) - 1
	c := p.Hand[idx]
	p.Hand = p.Hand[:idx]
	return c, nil
}

// AddCardsToBottom adds multiple cards to the bottom (front) of the hand
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

// GameState represents the current state of the game
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
