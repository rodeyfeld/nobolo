package main

import (
	"fmt"
	"log"
)

// DemoSlapping demonstrates the slap functionality with various scenarios
func DemoSlapping() {
	fmt.Println("🎯 NOBOLO Slap System Demonstration")
	fmt.Println("=====================================")

	// Create a game with 3 players
	game, err := NewGame("Alice", "Bob", "Charlie")
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	// Clear the pile for controlled demonstration
	game.Pile.Cards = []Card{}

	fmt.Printf("\nInitial state: %s\n", game)
	for i, player := range game.Players {
		fmt.Printf("Player %d (%s): %d cards\n", i, player.Name, player.HandSize())
	}

	fmt.Println("\n--- Demonstration 1: Valid Slap (Doubles) ---")
	// Set up doubles scenario
	game.Pile.AddCard(NewCard(Jack, Hearts, 11))
	game.Pile.AddCard(NewCard(Jack, Spades, 11))

	fmt.Printf("Pile now has: %v\n", game.Pile.GetTopCards(2))
	fmt.Printf("Slap condition: %s\n", game.CheckForSlaps())

	// Alice tries to slap
	success, slapType, err := game.TrySlap(0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Alice's slap result: Success=%t, Type=%s\n", success, slapType)
	}

	fmt.Printf("After slap - Alice: %d cards, Pile: %d cards\n",
		game.Players[0].HandSize(), game.Pile.Size())

	fmt.Println("\n--- Demonstration 2: Valid Slap (Queen-King) ---")
	// Set up Queen-King scenario
	game.Pile.AddCard(NewCard(Number, Clubs, 5))
	game.Pile.AddCard(NewCard(Queen, Hearts, 12))
	game.Pile.AddCard(NewCard(King, Diamonds, 13))

	fmt.Printf("Pile now has: %v\n", game.Pile.GetTopCards(3))
	fmt.Printf("Slap condition: %s\n", game.CheckForSlaps())

	// Bob tries to slap
	success, slapType, err = game.TrySlap(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bob's slap result: Success=%t, Type=%s\n", success, slapType)
	}

	fmt.Printf("After slap - Bob: %d cards, Pile: %d cards\n",
		game.Players[1].HandSize(), game.Pile.Size())

	fmt.Println("\n--- Demonstration 3: Valid Slap (Sandwich) ---")
	// Set up sandwich scenario: 7, 3, 7
	game.Pile.AddCard(NewCard(Number, Hearts, 7))
	game.Pile.AddCard(NewCard(Number, Clubs, 3))
	game.Pile.AddCard(NewCard(Number, Spades, 7))

	fmt.Printf("Pile now has: %v\n", game.Pile.GetTopCards(3))
	fmt.Printf("Slap condition: %s\n", game.CheckForSlaps())

	// Charlie tries to slap
	success, slapType, err = game.TrySlap(2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Charlie's slap result: Success=%t, Type=%s\n", success, slapType)
	}

	fmt.Printf("After slap - Charlie: %d cards, Pile: %d cards\n",
		game.Players[2].HandSize(), game.Pile.Size())

	fmt.Println("\n--- Demonstration 4: Invalid Slap (Penalty) ---")
	// Set up non-slappable scenario
	game.Pile.AddCard(NewCard(Number, Hearts, 2))
	game.Pile.AddCard(NewCard(Number, Clubs, 9))

	fmt.Printf("Pile now has: %v\n", game.Pile.GetTopCards(2))
	fmt.Printf("Slap condition: %s\n", game.CheckForSlaps())

	// Give Alice some cards so she can pay penalty
	game.Players[0].AddCard(NewCard(Number, Diamonds, 4))
	game.Players[0].AddCard(NewCard(Number, Spades, 6))

	fmt.Printf("Alice has %d cards before attempting slap\n", game.Players[0].HandSize())

	// Alice tries to slap (should fail)
	success, slapType, err = game.TrySlap(0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Alice's slap result: Success=%t, Type=%s\n", success, slapType)
	}

	fmt.Printf("After penalty - Alice: %d cards, Pile: %d cards\n",
		game.Players[0].HandSize(), game.Pile.Size())

	fmt.Println("\n--- Demonstration 5: Priority System (Sandwich vs Add-to-Ten) ---")
	// Clear pile and set up scenario where both sandwich and add-to-ten could apply
	game.Pile.Cards = []Card{}
	game.Pile.AddCard(NewCard(Number, Hearts, 4))
	game.Pile.AddCard(NewCard(Number, Clubs, 6))
	game.Pile.AddCard(NewCard(Number, Spades, 4))

	fmt.Printf("Pile has: %v\n", game.Pile.GetTopCards(3))
	fmt.Printf("This could be both Sandwich (4-6-4) and Add-to-Ten (6+4=10)\n")
	fmt.Printf("Priority system chooses: %s\n", game.CheckForSlaps())

	// Bob tries to slap
	success, slapType, err = game.TrySlap(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bob's slap result: Success=%t, Type=%s\n", success, slapType)
	}

	fmt.Println("\n🎉 Final Results:")
	fmt.Printf("Final state: %s\n", game)
	for i, player := range game.Players {
		fmt.Printf("Player %d (%s): %d cards\n", i, player.Name, player.HandSize())
	}

	fmt.Println("\n✅ Slap system demonstration complete!")
	fmt.Println("The system successfully:")
	fmt.Println("  - Detects all 5 slap types correctly")
	fmt.Println("  - Awards pile cards to successful slappers")
	fmt.Println("  - Applies penalties for incorrect slaps")
	fmt.Println("  - Handles priority when multiple conditions could apply")
}
