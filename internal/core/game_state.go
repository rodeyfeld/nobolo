package core

import "fmt"

type GameStatus byte

const (
	Unknown GameStatus = iota
	Running
	Over
)

type GameState struct {
	Players            []Player
	Pile               *Pile
	logLines           []string
	GameStatus         GameStatus
	ChallengeRemaining int
	ChallengeOwnerIdx  int // Index of the player who played the face card
}

func (gs *GameState) Log(msg string) {
	gs.logLines = append(gs.logLines, msg)
	fmt.Println(msg)
}

func (gs *GameState) Logs() []string {
	return gs.logLines
}
