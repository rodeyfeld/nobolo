package main

// GameState represents the current state of the game
type GameState byte

const (
	UnknownGameState = GameState(iota)
	GameStateGameRunning
	GameStateGameOver
)

// String returns the string representation of GameState
func (gs GameState) String() string {
	switch gs {
	case GameStateGameRunning:
		return "Running"
	case GameStateGameOver:
		return "Game Over"
	default:
		return "Unknown"
	}
}
