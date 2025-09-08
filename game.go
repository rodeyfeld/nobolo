package main

import (
	"fmt"
	"time"
)

// Pile represents the center pile of cards
type Pile struct {
	Cards []Card
}

// Game represents the entire game state
type Game struct {
	Players []*Player
	Pile    *Pile
	State   GameState
}

// AddCard adds a card to the pile
func (p *Pile) AddCard(card Card) {
	p.Cards = append(p.Cards, card)
}

// Size returns the number of cards in the pile
func (p *Pile) Size() int {
	return len(p.Cards)
}

// NewGame creates a new NOBOLO game with the specified players
func NewGame(playerNames ...string) (*Game, error) {
	if len(playerNames) < 2 {
		return nil, fmt.Errorf("need at least 2 players, got %d", len(playerNames))
	}

	deck := NewDeck()
	players := make([]*Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = NewPlayer(name)
	}

	deck.Shuffle()

	// Deal all cards in round-robin fashion
	playerIndex := 0
	for deck.Size() > 0 {
		card, err := deck.Draw()
		if err != nil {
			return nil, fmt.Errorf("error dealing cards: %w", err)
		}
		players[playerIndex].AddCard(card)
		playerIndex = (playerIndex + 1) % len(players)
	}

	return &Game{
		Players: players,
		Pile:    &Pile{Cards: make([]Card, 0)},
		State:   GameStateGameRunning,
	}, nil
}

// Play runs the main game loop
func (g *Game) Play() error {
	for g.State == GameStateGameRunning {
		for i := range g.Players {
			// Check win condition (player has all cards)
			if g.Players[i].HandSize() == 52 {
				fmt.Printf("%s wins! They got all the cards first.\n", g.Players[i].Name)
				g.State = GameStateGameOver
				return nil
			}

			// Player plays a card
			card, err := g.Players[i].PlayCard()
			if err != nil {
				return fmt.Errorf("player %s cannot play card: %w", g.Players[i].Name, err)
			}

			fmt.Printf("%s played %s\n", g.Players[i].Name, card)
			g.Pile.AddCard(card)

			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

// String returns a string representation of the game state
func (g *Game) String() string {
	return fmt.Sprintf("Game State: %s, Players: %d, Pile Size: %d",
		g.State, len(g.Players), g.Pile.Size())
}
