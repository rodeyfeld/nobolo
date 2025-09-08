package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("🎮 NOBOLO - Egyptian Rats Crew Card Game")
	fmt.Println("========================================")
	fmt.Println("Choose mode:")
	fmt.Println("1. Play original game (basic card playing)")
	fmt.Println("2. Slap system demonstration")
	fmt.Println()

	// For now, let's run the slap demo by default
	// In a real implementation, you'd get user input
	var choice int
	fmt.Print("Enter choice (1 or 2): ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		runOriginalGame()
	case 2:
		DemoSlapping()
	default:
		fmt.Println("Invalid choice, running slap demo...")
		DemoSlapping()
	}
}

func runOriginalGame() {
	fmt.Println("\n🎯 Starting Original NOBOLO Game")
	fmt.Println("=================================")

	// Create a new game with multiple players
	game, err := NewGame("Alice", "Bob", "Charlie")
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	fmt.Printf("Initial state: %s\n", game)

	// Print initial hand sizes
	for _, player := range game.Players {
		fmt.Printf("%s starts with %d cards\n", player.Name, player.HandSize())
	}
	fmt.Println()

	// Play the game
	if err := game.Play(); err != nil {
		log.Fatalf("Game error: %v", err)
	}

	fmt.Printf("Final state: %s\n", game)
}
