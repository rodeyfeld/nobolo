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

// AddCardsToBottom adds multiple cards to the bottom of the player's hand
// Used when a player wins a slap and gets the pile
func (p *Player) AddCardsToBottom(cards []Card) {
	// Add cards to the beginning of the hand (bottom of deck)
	p.Hand = append(cards, p.Hand...)
}

// RemoveTopCards removes n cards from the top of the player's hand
// Used for slap penalties - returns the removed cards
func (p *Player) RemoveTopCards(n int) ([]Card, error) {
	if n <= 0 {
		return []Card{}, nil
	}

	if len(p.Hand) == 0 {
		return []Card{}, nil
	}

	if len(p.Hand) < n {
		// Take all remaining cards if not enough
		cards := make([]Card, len(p.Hand))
		copy(cards, p.Hand)
		p.Hand = p.Hand[:0] // Clear the hand
		return cards, nil
	}

	// Remove from the end (top of deck)
	cards := make([]Card, n)
	copy(cards, p.Hand[len(p.Hand)-n:])
	p.Hand = p.Hand[:len(p.Hand)-n]

	return cards, nil
}
