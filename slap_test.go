package main

import (
	"testing"
)

func TestSlapDetection(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected SlapType
	}{
		{
			name: "Doubles - Two Jacks",
			cards: []Card{
				NewCard(Jack, Hearts, 11),
				NewCard(Jack, Spades, 11),
			},
			expected: Doubles,
		},
		{
			name: "Queen-King",
			cards: []Card{
				NewCard(Queen, Hearts, 12),
				NewCard(King, Spades, 13),
			},
			expected: QueenKing,
		},
		{
			name: "Add to Ten - 4 and 6",
			cards: []Card{
				NewCard(Number, Hearts, 4),
				NewCard(Number, Spades, 6),
			},
			expected: AddToTen,
		},
		{
			name: "Sandwich - 7, 3, 7",
			cards: []Card{
				NewCard(Number, Hearts, 7),
				NewCard(Number, Clubs, 3),
				NewCard(Number, Spades, 7),
			},
			expected: Sandwich,
		},
		{
			name: "Three in Order - 5, 6, 7",
			cards: []Card{
				NewCard(Number, Hearts, 5),
				NewCard(Number, Clubs, 6),
				NewCard(Number, Spades, 7),
			},
			expected: ThreeInOrder,
		},
		{
			name: "No Slap - Random cards",
			cards: []Card{
				NewCard(Number, Hearts, 2),
				NewCard(Number, Clubs, 9),
			},
			expected: NoSlap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckForSlap(tt.cards)
			if result != tt.expected {
				t.Errorf("CheckForSlap() = %v, want %v", result, tt.expected)
			}
		})
	}
}
