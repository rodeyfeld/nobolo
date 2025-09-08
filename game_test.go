package main

import (
	"testing"
)

func TestGameSlapping(t *testing.T) {
	// Create a game with 2 players
	game, err := NewGame("Alice", "Bob")
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	// Clear the pile and add specific cards to test slapping
	game.Pile.Cards = []Card{}

	t.Run("Valid Slap - Doubles", func(t *testing.T) {
		// Set up a doubles scenario
		game.Pile.AddCard(NewCard(Jack, Hearts, 11))
		game.Pile.AddCard(NewCard(Jack, Spades, 11))

		initialPileSize := game.Pile.Size()
		initialAliceCards := game.Players[0].HandSize()

		// Alice tries to slap
		success, slapType, err := game.TrySlap(0)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if !success {
			t.Error("Expected successful slap, got failure")
		}

		if slapType != Doubles {
			t.Errorf("Expected Doubles slap type, got %v", slapType)
		}

		// Check that Alice got the cards
		if game.Players[0].HandSize() != initialAliceCards+initialPileSize {
			t.Errorf("Alice should have gained %d cards, but hand size went from %d to %d",
				initialPileSize, initialAliceCards, game.Players[0].HandSize())
		}

		// Check that pile is empty
		if game.Pile.Size() != 0 {
			t.Errorf("Pile should be empty after successful slap, but has %d cards", game.Pile.Size())
		}
	})

	t.Run("Valid Slap - Queen-King", func(t *testing.T) {
		// Set up a Queen-King scenario
		game.Pile.AddCard(NewCard(Number, Clubs, 5))
		game.Pile.AddCard(NewCard(Queen, Hearts, 12))
		game.Pile.AddCard(NewCard(King, Spades, 13))

		initialPileSize := game.Pile.Size()
		initialBobCards := game.Players[1].HandSize()

		// Bob tries to slap
		success, slapType, err := game.TrySlap(1)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if !success {
			t.Error("Expected successful slap, got failure")
		}

		if slapType != QueenKing {
			t.Errorf("Expected QueenKing slap type, got %v", slapType)
		}

		// Check that Bob got the cards
		if game.Players[1].HandSize() != initialBobCards+initialPileSize {
			t.Errorf("Bob should have gained %d cards", initialPileSize)
		}

		// Check that pile is empty
		if game.Pile.Size() != 0 {
			t.Error("Pile should be empty after successful slap")
		}
	})

	t.Run("Valid Slap - Sandwich", func(t *testing.T) {
		// Set up a sandwich scenario: 7, 3, 7
		game.Pile.AddCard(NewCard(Number, Hearts, 7))
		game.Pile.AddCard(NewCard(Number, Clubs, 3))
		game.Pile.AddCard(NewCard(Number, Spades, 7))

		initialPileSize := game.Pile.Size()
		initialAliceCards := game.Players[0].HandSize()

		// Alice tries to slap
		success, slapType, err := game.TrySlap(0)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if !success {
			t.Error("Expected successful slap, got failure")
		}

		if slapType != Sandwich {
			t.Errorf("Expected Sandwich slap type, got %v", slapType)
		}

		// Check that Alice got the cards
		if game.Players[0].HandSize() != initialAliceCards+initialPileSize {
			t.Errorf("Alice should have gained %d cards", initialPileSize)
		}
	})

	t.Run("Invalid Slap - Penalty", func(t *testing.T) {
		// Set up a non-slappable scenario
		game.Pile.AddCard(NewCard(Number, Hearts, 2))
		game.Pile.AddCard(NewCard(Number, Clubs, 9))

		initialPileSize := game.Pile.Size()
		initialBobCards := game.Players[1].HandSize()

		// Bob tries to slap (should fail)
		success, slapType, err := game.TrySlap(1)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if success {
			t.Error("Expected failed slap, got success")
		}

		if slapType != NoSlap {
			t.Errorf("Expected NoSlap type, got %v", slapType)
		}

		// Check that Bob lost 2 cards (penalty)
		expectedBobCards := initialBobCards - 2
		if expectedBobCards < 0 {
			expectedBobCards = 0 // Can't go below 0
		}

		if game.Players[1].HandSize() != expectedBobCards {
			t.Errorf("Bob should have lost 2 cards, hand size went from %d to %d",
				initialBobCards, game.Players[1].HandSize())
		}

		// Check that pile gained penalty cards
		expectedPileSize := initialPileSize + 2
		if initialBobCards < 2 {
			expectedPileSize = initialPileSize + initialBobCards
		}

		if game.Pile.Size() != expectedPileSize {
			t.Errorf("Pile should have gained penalty cards, expected %d, got %d",
				expectedPileSize, game.Pile.Size())
		}
	})
}

func TestGameSlapPriority(t *testing.T) {
	game, err := NewGame("Alice", "Bob")
	if err != nil {
		t.Fatalf("Failed to create game: %v", err)
	}

	t.Run("Sandwich vs Add-to-Ten Priority", func(t *testing.T) {
		// Clear pile and set up a scenario where both sandwich and add-to-ten could apply
		// 4, 6, 4 - this is both a sandwich (4-6-4) and the last two add to 10 (6+4)
		// Sandwich should have higher priority
		game.Pile.Cards = []Card{}
		game.Pile.AddCard(NewCard(Number, Hearts, 4))
		game.Pile.AddCard(NewCard(Number, Clubs, 6))
		game.Pile.AddCard(NewCard(Number, Spades, 4))

		slapType := game.CheckForSlaps()
		if slapType != Sandwich {
			t.Errorf("Expected Sandwich to have priority over Add-to-Ten, got %v", slapType)
		}

		// Test the slap
		success, resultType, err := game.TrySlap(0)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if !success {
			t.Error("Expected successful slap")
		}

		if resultType != Sandwich {
			t.Errorf("Expected Sandwich slap result, got %v", resultType)
		}
	})
}

func TestGameSlapIntegration(t *testing.T) {
	t.Run("Complete Slap Scenario", func(t *testing.T) {
		// Create a game and manually set up a scenario
		game, err := NewGame("Alice", "Bob", "Charlie")
		if err != nil {
			t.Fatalf("Failed to create game: %v", err)
		}

		// Clear all hands and pile for controlled testing
		for _, player := range game.Players {
			player.Hand = []Card{}
		}
		game.Pile.Cards = []Card{}

		// Give each player some cards
		game.Players[0].AddCard(NewCard(Number, Hearts, 2))
		game.Players[0].AddCard(NewCard(Number, Clubs, 3))
		game.Players[1].AddCard(NewCard(Number, Spades, 4))
		game.Players[1].AddCard(NewCard(Number, Diamonds, 5))
		game.Players[2].AddCard(NewCard(Jack, Hearts, 11))

		// Simulate a game sequence that leads to a slappable condition
		// Alice plays 2
		card1, _ := game.Players[0].PlayCard()
		game.Pile.AddCard(card1)

		// Bob plays 4
		card2, _ := game.Players[1].PlayCard()
		game.Pile.AddCard(card2)

		// Charlie plays Jack
		card3, _ := game.Players[2].PlayCard()
		game.Pile.AddCard(card3)

		// Alice plays 3
		card4, _ := game.Players[0].PlayCard()
		game.Pile.AddCard(card4)

		// Bob plays 5
		card5, _ := game.Players[1].PlayCard()
		game.Pile.AddCard(card5)

		// Now add Queen then King to create Queen-King slap condition
		game.Pile.AddCard(NewCard(Queen, Spades, 12))
		game.Pile.AddCard(NewCard(King, Diamonds, 13))

		// Check what's in the pile
		topCards := game.Pile.GetTopCards(3)
		t.Logf("Top 3 cards in pile: %v", topCards)

		// Check that slap condition exists
		slapType := game.CheckForSlaps()
		if slapType != QueenKing {
			t.Errorf("Expected QueenKing slap condition, got %v. Pile has %d cards", slapType, game.Pile.Size())
		}

		// Alice slaps and should get all the cards
		initialAliceCards := game.Players[0].HandSize()
		initialPileSize := game.Pile.Size()

		success, resultType, err := game.TrySlap(0)
		if err != nil {
			t.Fatalf("Error during slap: %v", err)
		}

		if !success {
			t.Error("Expected successful slap")
		}

		if resultType != QueenKing {
			t.Errorf("Expected QueenKing result, got %v", resultType)
		}

		// Verify Alice got all the pile cards
		if game.Players[0].HandSize() != initialAliceCards+initialPileSize {
			t.Errorf("Alice should have %d cards, but has %d",
				initialAliceCards+initialPileSize, game.Players[0].HandSize())
		}

		// Verify pile is empty
		if game.Pile.Size() != 0 {
			t.Errorf("Pile should be empty, but has %d cards", game.Pile.Size())
		}
	})
}
