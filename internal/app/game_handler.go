package app

import (
	"fmt"
	"image/color"

	"nobolo/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Players          []core.Player
	Pile             *core.Pile
	CurrentPlayer    int
	challengeOwner   int
	remainingChances int
	logLines         []string
}

type GameHandler struct {
	Game      *Game
	GameState core.GameState
	logLines  []string
}

func NewGameHandler() *GameHandler {
	gh := &GameHandler{
		GameState: core.UnknownGameState,
		Game:      &Game{},
		logLines:  make([]string, 0, 64),
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

	for deck.CardCount() > 0 {
		for i := range g.Players {
			card, err := deck.Draw()
			if err != nil {
				break
			}
			g.Players[i].AddCardsToBottom([]core.Card{card})
		}
	}

	g.challengeOwner = -1
	g.remainingChances = 0
	g.logLines = g.logLines[:0]
	g.appendLog("Game started")
}

// Draw implements ebiten.Game
func (gh *GameHandler) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{24, 28, 36, 255})

	if gh.GameState == core.UnknownGameState || gh.GameState == core.GameStateGameOver {
		drawMenu(screen)
	} else if gh.GameState == core.GameStateGameRunning {
		drawGame(screen, gh.Game)
	}
}

// Layout implements ebiten.Game
func (gh *GameHandler) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func drawMenu(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "🎮 NOBOLO - Egyptian Rats Crew Card Game", 200, 100)
	ebitenutil.DebugPrintAt(screen, "[Click] Play Game", 350, 225)
}

func drawGame(screen *ebiten.Image, g *Game) {
	ebitenutil.DebugPrintAt(screen, "Game in Progress", 50, 50)

	yPos := 100
	for i, player := range g.Players {
		text := fmt.Sprintf("%s: %d cards", player.Name, len(player.Hand))
		if i == g.CurrentPlayer {
			text += " (Your turn)"
		}
		ebitenutil.DebugPrintAt(screen, text, 50, yPos)
		yPos += 30
	}

	pileText := fmt.Sprintf("Pile: %d cards", len(g.Pile.Cards))
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

// checkWinCondition checks if only one player remains with cards
func (g *Game) checkWinCondition() bool {

	countWithCards := 0
	lastIdx := -1
	for i := range g.Players {
		if len(g.Players[i].Hand) > 0 {
			countWithCards++
			lastIdx = i
		}
	}

	if countWithCards <= 1 {
		// Show winner inline by moving turn to winner if exists
		if lastIdx >= 0 {
			g.CurrentPlayer = lastIdx
		}
		return true
	}

	return false
}

func (g *Game) appendLog(s string) {
	g.logLines = append(g.logLines, s)
	fmt.Println(s)
}
