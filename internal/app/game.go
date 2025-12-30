package app

import (
	"fmt"
	"image/color"

	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SimpleGame struct {
	GameState     core.GameState
	Players       []core.Player
	Pile          []Card 
	CurrentPlayer int
	challengeOwner   int
	remainingChances int
	// UI log
	logLines []string
}

func NewSimpleGame() *SimpleGame {
	g := &SimpleGame{
		GameState:        core.UnknownGameState,
		Players:          make([]core.Player, 0),
		Pile:             make([]core.Card, 0),
		challengeOwner:   -1,
		remainingChances: 0,
		logLines:         make([]string, 0, 64),
	}
	g.startGame()
	return g
}


func (g *Simplegame) createPlayers() { 
	playerNames := []string{"Alice", "Bob", "Charlie"}
	for _, name := range playerNames {
		g.Players = append(g.Players, *core.NewPlayer(name))
	}
}

func (g *SimpleGame) startGame() {
	// Reset game state
	g.GameState = core.GameStateGameRunning
	g.CurrentPlayer = 0
	players = g.createPlayers()

	// Deal cards
	deck := core.NewDeck()
	deck.Shuffle()

	for len(deck.cards) > 0 {
		for player := range players {
			card, err := deck.Draw()
			if err != nil {
				break
			}
			g.Players[playerIndex].AddCard(card)
		}
	}

	g.challengeOwner = -1
	g.remainingChances = 0
	g.logLines = g.logLines[:0]
	g.appendLog("Game started")
}

func (g *SimpleGame) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.GameState == core.UnknownGameState || g.GameState == core.GameStateGameOver {
			if x >= 350 && x <= 550 && y >= 200 && y <= 250 {
				g.startGame()
			}
		}
	}
	g.GameLoop()
	return nil
}

// Draw implements ebiten.Game
func (g *SimpleGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{24, 28, 36, 255})

	if g.GameState == core.UnknownGameState || g.GameState == core.GameStateGameOver {
		drawMenu(screen)
	} else if g.GameState == core.GameStateGameRunning {
		drawGame(screen, g)
	}
}

// Layout implements ebiten.Game
func (g *SimpleGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "🎮 NOBOLO - Egyptian Rats Crew Card Game", 200, 100)
	ebitenutil.DebugPrintAt(screen, "[Click] Play Game", 350, 225)
}

func drawGame(screen *ebiten.Image, g *SimpleGame) {
	ebitenutil.DebugPrintAt(screen, "Game in Progress", 50, 50)

	yPos := 100
	for i, player := range g.Players {
		text := fmt.Sprintf("%s: %d cards", player.Name, player.HandSize())
		if i == g.CurrentPlayer && g.GameState == core.GameStateGameRunning {
			text += " (Your turn)"
		}
		ebitenutil.DebugPrintAt(screen, text, 50, yPos)
		yPos += 30
	}

	pileText := fmt.Sprintf("Pile: %d cards", g.Pile.Size())
	ebitenutil.DebugPrintAt(screen, pileText, 50, yPos+50)

	stateText := fmt.Sprintf("Game State: %s", g.GameState.String())
	ebitenutil.DebugPrintAt(screen, stateText, 50, 550)

	// Controls hint
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

// checkWinCondition checks if only one player remains with cards
func (g *SimpleGame) checkWinCondition() bool {
	if g.GameState != core.GameStateGameRunning {
		return false
	}

	countWithCards := 0
	lastIdx := -1
	for i := range g.Players {
		if len(g.Players[i].hand) > 0 {
			countWithCards++
			lastIdx = i
		}
	}

	if countWithCards <= 1 {
		g.GameState = core.GameStateGameOver
		// Show winner inline by moving turn to winner if exists
		if lastIdx >= 0 {
			g.CurrentPlayer = lastIdx
		}
		return true
	}

	return false
}

func (g *SimpleGame) appendLog(s string) {
	g.logLines = append(g.logLines, s)
	fmt.Println(s)
}
