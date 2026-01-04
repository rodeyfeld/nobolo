package core

import "errors"

// Turn represents a single snapshot in the game history
type Turn struct {
	GameState   *GameState
	Player      *Player
	NextPlayer *Player
}

// TurnNode is a node in the TurnHistory linked list
type TurnNode struct {
	Value Turn
	Next  *TurnNode
	Prev  *TurnNode
}

// TurnHistory manages a linked list of game turns (History)
type TurnHistory struct {
	Head  *TurnNode
	Tail  *TurnNode
	count int
}

func (th *TurnHistory) Count() int {
	return th.count
}

// Push adds a turn to the end of the history
func (th *TurnHistory) Push(turn Turn) {
	node := &TurnNode{Value: turn}
	if th.Tail == nil {
		th.Head = node
		th.Tail = node
	} else {
		node.Prev = th.Tail
		th.Tail.Next = node
		th.Tail = node
	}
	th.count++
}

// Pop removes and returns the most recent turn
func (th *TurnHistory) Pop() (Turn, error) {
	if th.Tail == nil {
		return Turn{}, errors.New("history is empty")
	}

	val := th.Tail.Value
	th.Tail = th.Tail.Prev

	if th.Tail == nil {
		th.Head = nil
	} else {
		th.Tail.Next = nil
	}

	th.count--
	return val, nil
}

// Latest returns the most recent turn without removing it
func (th *TurnHistory) Latest() (*Turn, error) {
	if th.Tail == nil {
		return nil, errors.New("history is empty")
	}
	return &th.Tail.Value, nil
}
