package main

import (
	"fmt"
	"os"
	"strconv"

	"nobolo/internal/app"
)

func main() {
	fmt.Println("🎮 NOBOLO - Egyptian Rats Crew Card Game")
	fmt.Println("========================================")
	fmt.Println("Choose mode:")
	fmt.Println("1. Simple ECS-based game with UI")
	fmt.Println()

	// Check if choice was provided as command line argument
	var choice int = 1 // Default to ECS game
	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			choice = val
		}
	} else {
		// Interactive mode - get user input
		fmt.Print("Enter choice (1): ")
		fmt.Scanf("%d", &choice)
	}

	switch choice {
	case 1:
		app.Run()
	default:
		fmt.Println("Running ECS game...")
		app.Run()
	}
}
