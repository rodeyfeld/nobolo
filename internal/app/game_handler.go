package app

import (
	"fmt"
	"image/color"
	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var faceChances = map[core.CardFace]int{
	core.Jack:  1,
	core.Queen: 2,
	core.King:  3,
	core.Ace:   4,
}

type Game struct {
	InitialGameState *core.GameState
	History          *core.TurnHistory
	CurrentTurn      *core.Turn
}

type GameHandler struct {
	Game       *Game
	GameStatus core.GameStatus
}

func NewGameHandler() *GameHandler {
	players := createPlayers()
	initialGameState := &core.GameState{
		Players:           players,
		Pile:              &core.Pile{},
		ChallengeOwnerIdx: -1,
	}

	game := &Game{
		InitialGameState: initialGameState,
		History:          &core.TurnHistory{},
	}

	gh := &GameHandler{
		GameStatus: core.Unknown,
		Game:       game,
	}
	return gh
}

func createPlayers() []core.Player {
	playerNames := []string{"Alice", "Bob", "Charlie"}
	players := make([]core.Player, 0, len(playerNames))
	for _, name := range playerNames {
		players = append(players, *core.NewPlayer(name))
	}
	return players
}
func (g *Game) isGameWon() bool {
	aliveCount := 0
	for _, p := range g.InitialGameState.Players {
		if p.Alive {
			aliveCount++
		}
	}
	return aliveCount <= 1
}

func (g *Game) appendLog(msg string) {
	if g.CurrentTurn != nil {
		g.CurrentTurn.GameState.Log(msg)
	}
}

func (g *Game) startGame() {
	// 1. Setup Initial State
	state := g.InitialGameState
	state.ChallengeOwnerIdx = -1
	state.ChallengeRemaining = 0
	for i := range state.Players {
		state.Players[i].Alive = true
		state.Players[i].Hand.Clear()
	}

	// 2. Deal cards
	deck := core.NewDeck()
	deck.Shuffle()

	for deck.Count() > 0 {
		for i := range state.Players {
			card, err := deck.Draw()
			if err != nil {
				break
			}
			state.Players[i].Hand.PushBottom([]core.Card{card})
		}
	}

	// 3. Create First Turn
	g.CurrentTurn = &core.Turn{
		GameState:  state,
		Player:     &state.Players[0],
		NextPlayer: &state.Players[1],
	}

	g.History = &core.TurnHistory{}
}

// Update implements ebiten.Game
func (gh *GameHandler) Update() error {
	if gh.GameStatus == core.Running {
		// Playing a card
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gh.Game.ResolveCurrentTurn()
		}

		// Slap attempts for each player
		// Alice: A, Bob: S, Charlie: D
		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			gh.Game.HandleSlap(0)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			gh.Game.HandleSlap(1)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			gh.Game.HandleSlap(2)
		}

		// Win check every tick
		if gh.Game.isGameWon() {
			gh.GameStatus = core.Over
		}
	} else if (gh.GameStatus == core.Unknown || gh.GameStatus == core.Over) &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gh.GameStatus = core.Running
		gh.Game.startGame()
	}
	return nil
}

// Draw implements ebiten.Game
func (gh *GameHandler) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{24, 28, 36, 255})

	if gh.GameStatus == core.Unknown || gh.GameStatus == core.Over {
		gh.drawMenu(screen)
	} else {
		gh.drawGame(screen)
	}
}

// Layout implements ebiten.Game
func (gh *GameHandler) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (gh *GameHandler) drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "🎮 NOBOLO - Egyptian Rats Crew", 300, 250)
	ebitenutil.DebugPrintAt(screen, "Click to Start", 350, 300)
}

func (gh *GameHandler) drawGame(screen *ebiten.Image) {
	if gh.Game.CurrentTurn == nil {
		return
	}

	state := gh.Game.CurrentTurn.GameState
	y := 50
	ebitenutil.DebugPrintAt(screen, "Players:", 50, y)
	y += 20
	for i := range state.Players {
		p := &state.Players[i]
		status := ""
		if !p.Alive {
			status = " (OUT)"
		} else if gh.Game.CurrentTurn.Player == p {
			status = " <--- Current Turn"
		}
		text := fmt.Sprintf("%s: %d cards%s", p.Name, p.Hand.Count(), status)
		ebitenutil.DebugPrintAt(screen, text, 50, y)
		y += 20
	}

	y += 20
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Pile: %d cards", state.Pile.Count()), 50, y)

	if state.ChallengeRemaining > 0 {
		y += 20
		owner := state.Players[state.ChallengeOwnerIdx].Name
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("CHALLENGE: %s (chances left: %d)", owner, state.ChallengeRemaining), 50, y)
	}

	// Draw Logs
	y = 50
	ebitenutil.DebugPrintAt(screen, "Events:", 500, y)
	y += 20
	logs := state.Logs()
	start := 0
	if len(logs) > 20 {
		start = len(logs) - 20
	}
	for _, line := range logs[start:] {
		ebitenutil.DebugPrintAt(screen, line, 500, y)
		y += 15
	}

	ebitenutil.DebugPrintAt(screen, "SPACE: Play | Slap -> Alice: A, Bob: S, Charlie: D", 250, 550)
}
