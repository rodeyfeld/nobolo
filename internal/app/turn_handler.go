package app

import (
	"fmt"
	"nobolo/internal/core"
)

func (g *Game) ResolveCurrentTurn() {
	if g.CurrentTurn == nil {
		return
	}

	turn := g.CurrentTurn
	state := turn.GameState
	player := turn.Player

	// 1. Check if player has cards
	if player.Hand.Count() == 0 {
		g.appendLog(fmt.Sprintf("%s is out of cards!", player.Name))
		player.Alive = false

		if state.ChallengeRemaining > 0 {
			// If they were in a challenge and ran out of cards, the challenge owner wins
			winner := &state.Players[state.ChallengeOwnerIdx]
			cards := state.Pile.Clear()
			winner.Hand.PushBottom(cards)
			g.appendLog(fmt.Sprintf("%s ran out of cards during challenge! %s takes the pile", player.Name, winner.Name))
			state.ChallengeRemaining = 0
			state.ChallengeOwnerIdx = -1
			g.setNextTurn(winner)
		} else {
			g.moveToNextTurn()
		}
		return
	}

	// 2. Player plays a card
	card, _ := player.Hand.PopTop()
	state.Pile.PushTop(card)
	g.appendLog(fmt.Sprintf("%s played %s", player.Name, formatCard(card)))

	// 3. Check for face card (Challenge)
	if chances, isFace := faceChances[card.Face]; isFace {
		state.ChallengeRemaining = chances
		state.ChallengeOwnerIdx = g.getPlayerIndex(player)
		g.appendLog(fmt.Sprintf("CHALLENGE! %s has %d chances", state.Players[(g.getPlayerIndex(player)+1)%len(state.Players)].Name, chances))
		g.moveToNextTurn()
		return
	}

	// 4. If in a challenge, decrement chances
	if state.ChallengeRemaining > 0 {
		state.ChallengeRemaining--
		if state.ChallengeRemaining == 0 {
			// Challenge failed! Previous face card player takes the pile
			winner := &state.Players[state.ChallengeOwnerIdx]
			cards := state.Pile.Clear()
			winner.Hand.PushBottom(cards)
			g.appendLog(fmt.Sprintf("%s failed challenge! %s takes the pile", player.Name, winner.Name))

			// Reset challenge
			state.ChallengeRemaining = 0
			state.ChallengeOwnerIdx = -1

			// Winner of challenge goes next
			g.setNextTurn(winner)
			return
		}
		// Same player keeps playing, but we want a new Turn object for each card
		g.setNextTurn(player)
		return
	}

	// 5. Normal play, move to next player
	g.moveToNextTurn()
}

func (g *Game) HandleSlap(playerIdx int) {
	if g.CurrentTurn == nil {
		return
	}
	state := g.CurrentTurn.GameState
	if playerIdx < 0 || playerIdx >= len(state.Players) {
		return
	}
	player := &state.Players[playerIdx]

	slapType := core.CheckSlapType(state.Pile)
	if slapType != core.NoSlap {
		count := state.Pile.Count()
		cards := state.Pile.Clear()
		player.Hand.PushBottom(cards)
		g.appendLog(fmt.Sprintf("%s SLAPPED! %s (+%d cards)", player.Name, slapType.String(), count))

		// If a player was out and slaps back in, they are alive again
		if !player.Alive {
			player.Alive = true
			g.appendLog(fmt.Sprintf("%s slaps back into the game!", player.Name))
		}

		// Reset challenge if any
		state.ChallengeRemaining = 0
		state.ChallengeOwnerIdx = -1

		// Slaver goes next
		g.setNextTurn(player)
	} else {
		// Bad slap - only penalize if they have something to lose
		if !player.Alive {
			return
		}

		penalty := player.Hand.PopTopN(2)
		state.Pile.PushBottom(penalty)
		g.appendLog(fmt.Sprintf("%s MISIDENTIFIED slap! (-%d cards)", player.Name, len(penalty)))

		// Check if penalty made them lose
		if player.Hand.Count() == 0 {
			player.Alive = false
			g.appendLog(fmt.Sprintf("%s ran out of cards from penalty!", player.Name))
			if g.CurrentTurn.Player == player {
				g.moveToNextTurn()
			}
		}
	}
}

func (g *Game) getPlayerIndex(p *core.Player) int {
	for i := range g.InitialGameState.Players {
		if &g.InitialGameState.Players[i] == p {
			return i
		}
	}
	return -1
}

func (g *Game) moveToNextTurn() {
	state := g.CurrentTurn.GameState
	currIdx := g.getPlayerIndex(g.CurrentTurn.Player)

	// Find next alive player
	nextIdx := (currIdx + 1) % len(state.Players)
	for !state.Players[nextIdx].Alive {
		nextIdx = (nextIdx + 1) % len(state.Players)
		if nextIdx == currIdx {
			break // Only one player alive
		}
	}

	g.setNextTurn(&state.Players[nextIdx])
}

func (g *Game) setNextTurn(nextPlayer *core.Player) {
	// Push current turn to history
	if g.CurrentTurn != nil {
		g.History.Push(*g.CurrentTurn)
	}

	// Create new turn
	state := g.InitialGameState
	currIdx := g.getPlayerIndex(nextPlayer)
	nextIdx := (currIdx + 1) % len(state.Players)
	for !state.Players[nextIdx].Alive {
		nextIdx = (nextIdx + 1) % len(state.Players)
	}

	g.CurrentTurn = &core.Turn{
		GameState:  state,
		Player:     nextPlayer,
		NextPlayer: &state.Players[nextIdx],
	}
}
