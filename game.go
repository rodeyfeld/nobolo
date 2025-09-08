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

// CheckSlap checks if the current pile state allows for a valid slap
func (p *Pile) CheckSlap() SlapType {
	return CheckForSlap(p.Cards)
}

// TakeAllCards removes all cards from the pile and returns them
// Used when a player successfully slaps the pile
func (p *Pile) TakeAllCards() []Card {
	cards := make([]Card, len(p.Cards))
	copy(cards, p.Cards)
	p.Cards = p.Cards[:0] // Clear the pile
	return cards
}

// GetTopCards returns the top n cards from the pile without removing them
// Useful for displaying recent cards or checking slap conditions
func (p *Pile) GetTopCards(n int) []Card {
	if n <= 0 || len(p.Cards) == 0 {
		return []Card{}
	}

	start := len(p.Cards) - n
	if start < 0 {
		start = 0
	}

	return p.Cards[start:]
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

// TrySlap attempts to slap the pile for the given player
// Returns true if the slap was successful, false if it was a penalty
func (g *Game) TrySlap(playerIndex int) (bool, SlapType, error) {
	if playerIndex < 0 || playerIndex >= len(g.Players) {
		return false, NoSlap, fmt.Errorf("invalid player index: %d", playerIndex)
	}

	player := g.Players[playerIndex]
	slapType := g.Pile.CheckSlap()

	if slapType != NoSlap {
		// Valid slap! Player gets all the cards
		cards := g.Pile.TakeAllCards()
		player.AddCardsToBottom(cards)
		fmt.Printf("%s successfully slapped for %s! Got %d cards.\n",
			player.Name, slapType, len(cards))
		return true, slapType, nil
	} else {
		// Invalid slap! Player loses 2 cards as penalty
		penaltyCards, err := player.RemoveTopCards(2)
		if err != nil {
			return false, NoSlap, fmt.Errorf("error applying penalty: %w", err)
		}

		// Add penalty cards to bottom of pile
		for _, card := range penaltyCards {
			g.Pile.AddCard(card)
		}

		fmt.Printf("%s slapped incorrectly! Lost %d cards as penalty.\n",
			player.Name, len(penaltyCards))
		return false, NoSlap, nil
	}
}

// CheckForSlaps checks if any slap conditions are met after a card is played
func (g *Game) CheckForSlaps() SlapType {
	return g.Pile.CheckSlap()
}
