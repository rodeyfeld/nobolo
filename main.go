package main

import (
	"fmt"
	"log"
)

func main() {
	// Create a new game with multiple players
	game, err := NewGame("Alice", "Bob", "Charlie")
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	fmt.Printf("Starting NOBOLO game!\n")
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
