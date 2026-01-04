package core

import (
	"errors"
	"math/rand/v2"
)

// CardStack provides common functionality for managing a collection of cards
type CardStack struct {
	Cards []Card
}

func (s *CardStack) Count() int {
	return len(s.Cards)
}

func (s *CardStack) Shuffle() {
	rand.Shuffle(len(s.Cards), func(i, j int) {
		s.Cards[i], s.Cards[j] = s.Cards[j], s.Cards[i]
	})
}

func (s *CardStack) PushTop(card Card) {
	s.Cards = append(s.Cards, card)
}

func (s *CardStack) PushBottom(cards []Card) {
	// Prepend cards to the bottom (index 0)
	s.Cards = append(cards, s.Cards...)
}

func (s *CardStack) PopTop() (Card, error) {
	if len(s.Cards) == 0 {
		return Card{}, errors.New("stack is empty")
	}
	idx := len(s.Cards) - 1
	c := s.Cards[idx]
	s.Cards = s.Cards[:idx]
	return c, nil
}

// PopTopN removes and returns n cards from the top
func (s *CardStack) PopTopN(n int) []Card {
	if n <= 0 || len(s.Cards) == 0 {
		return []Card{}
	}
	if len(s.Cards) < n {
		n = len(s.Cards)
	}
	start := len(s.Cards) - n
	cards := make([]Card, n)
	copy(cards, s.Cards[start:])
	s.Cards = s.Cards[:start]
	return cards
}

func (s *CardStack) Clear() []Card {
	cards := make([]Card, len(s.Cards))
	copy(cards, s.Cards)
	s.Cards = s.Cards[:0]
	return cards
}



