package main

import "fmt"

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

// String returns the string representation of CardFace
func (f CardFace) String() string {
	switch f {
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	case Number:
		return "Number"
	default:
		return "Unknown"
	}
}

// String returns the string representation of CardSuit
func (s CardSuit) String() string {
	switch s {
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

// Card represents a single playing card
type Card struct {
	Face  CardFace
	Suit  CardSuit
	Value int
}

// String returns the string representation of a card
func (c Card) String() string {
	if c.Face == Number {
		return fmt.Sprintf("%d of %s", c.Value, c.Suit)
	}
	return fmt.Sprintf("%s of %s", c.Face, c.Suit)
}

// NewCard creates a new card with the given face, suit, and value
func NewCard(face CardFace, suit CardSuit, value int) Card {
	return Card{
		Face:  face,
		Suit:  suit,
		Value: value,
	}
}

// IsRed returns true if the card is red (Hearts or Diamonds)
func (c Card) IsRed() bool {
	return c.Suit == Hearts || c.Suit == Diamonds
}

// IsBlack returns true if the card is black (Clubs or Spades)
func (c Card) IsBlack() bool {
	return c.Suit == Clubs || c.Suit == Spades
}

// IsFaceCard returns true if the card is a face card (Jack, Queen, King, or Ace)
func (c Card) IsFaceCard() bool {
	return c.Face == Jack || c.Face == Queen || c.Face == King || c.Face == Ace
}
