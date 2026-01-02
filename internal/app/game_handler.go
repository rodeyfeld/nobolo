package app

import (
	"fmt"

	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var faceChances = map[core.CardFace]int{
	core.Jack:  1,
	core.Queen: 2,
	core.King:  3,
	core.Ace:   4,
}

type Game struct {
	Players          []core.Player
	Pile             *core.Pile
	CurrentPlayer    int
	challengeOwner   int
	remainingChances int
	logLines         []string
	Turns *core.Turn
}

type GameHandler struct {
	Game      *Game
	GameState core.GameState
	logLines  []string
}

func NewGameHandler() *GameHandler {
	gh := &GameHandler{
		GameState: core.UnknownGameState,
		Game: &Game{
			Pile: &core.Pile{},
		},
		logLines: make([]string, 0, 64),
	}
	return gh
}

func (g *Game) createPlayers() {
	playerNames := []string{"Alice", "Bob", "Charlie"}
	for _, name := range playerNames {
		g.Players = append(g.Players, *core.NewPlayer(name))
	}
}

func (g *Game) startGame() {
	// Reset game state
	g.CurrentPlayer = 0
	g.createPlayers()

	// Deal cards
	deck := core.NewDeck()
	deck.Shuffle()

	for deck.Count() > 0 {
		for i := range g.Players {
			card, err := deck.Draw()
			if err != nil {
				break
			}
			g.Players[i].Hand.PushBottom([]core.Card{card})
		}
	}

	g.challengeOwner = -1
	g.remainingChances = 0
	g.logLines = g.logLines[:0]
	g.appendLog("Game started")
}

func drawGame(screen *ebiten.Image, g *Game) {
	ebitenutil.DebugPrintAt(screen, "Game in Progress", 50, 50)

	yPos := 100
	for i, player := range g.Players {
		text := fmt.Sprintf("%s: %d cards", player.Name, player.Hand.Count())
		if i == g.CurrentPlayer {
			text += " (Your turn)"
		}
		ebitenutil.DebugPrintAt(screen, text, 50, yPos)
		yPos += 30
	}

	pileText := fmt.Sprintf("Pile: %d cards", g.Pile.Count())
	ebitenutil.DebugPrintAt(screen, pileText, 50, yPos+50)

	ebitenutil.DebugPrintAt(screen, "Controls: SPACE=play, S=slap, Click=Start/Restart", 400, 550)

	// Draw log on the right
	y := 80
	ebitenutil.DebugPrintAt(screen, "Events:", 500, y)
	y += 20
	start := 0
	if len(g.logLines) > 18 {
		start = len(g.logLines) - 18
	}
	for _, line := range g.logLines[start:] {
		ebitenutil.DebugPrintAt(screen, line, 500, y)
		y += 16
	}
}

func (g *Game) isGameWon() bool {
	alivePlayerCount := 0
	for _, player := range g.Players {
		if player.Alive {

			alivePlayerCount += 1
		}
	}
	if alivePlayerCount > 1 {
		return false
	}
	return true
}
