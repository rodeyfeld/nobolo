package main

import (
	"fmt"
	"math/rand/v2"
)

// Deck represents a collection of cards
type Deck struct {
	cards []Card
}

// Card definitions using a more declarative approach
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
			cards = append(cards, NewCard(def.face, suit, def.value))
		}
	}

	return &Deck{cards: cards}
}

// Shuffle randomizes the order of cards in the deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Draw removes and returns the top card from the deck
func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, fmt.Errorf("deck is empty")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

// Size returns the number of cards remaining in the deck
func (d *Deck) Size() int {
	return len(d.cards)
}
