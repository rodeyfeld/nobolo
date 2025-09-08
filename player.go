package main

import "errors"

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

// PlayCard removes and returns the last card from the player's hand
func (p *Player) PlayCard() (Card, error) {
	if len(p.Hand) == 0 {
		return Card{}, errors.New("player has no cards")
	}

	lastIndex := len(p.Hand) - 1
	cardToPlay := p.Hand[lastIndex]
	p.Hand = p.Hand[:lastIndex]

	return cardToPlay, nil
}

// AddCard adds a card to the player's hand
func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// HandSize returns the number of cards in the player's hand
func (p *Player) HandSize() int {
	return len(p.Hand)
}
