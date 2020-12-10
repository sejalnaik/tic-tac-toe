package game

import (
	"github.com/sejalnaik/tic-tac-toe/analyzer"
	"github.com/sejalnaik/tic-tac-toe/board"
	gamestatus "github.com/sejalnaik/tic-tac-toe/gameStatus"
	"github.com/sejalnaik/tic-tac-toe/player"
)

// Game is to create game type
type Game struct {
	players       []*player.Player
	board         *board.Board
	analyzer      *analyzer.Analyzer
	gameStatus    string
	addToken      bool
	currentPlayer *player.Player
	nextPlayer    *player.Player
}

// NewGame is to create a new struct of type game
func NewGame(playersList []*player.Player, size int) *Game {
	tempBoard := board.NewBoard(size)
	tempAnalyzer := analyzer.NewAnalyzer(tempBoard)
	tempGame := &Game{
		players:       playersList,
		board:         tempBoard,
		analyzer:      tempAnalyzer,
		gameStatus:    gamestatus.InProgress,
		addToken:      true,
		currentPlayer: playersList[0],
		nextPlayer:    playersList[1],
	}
	return tempGame
}

// IsAddToken to chcek if cell is marked
func (g *Game) IsAddToken() bool {
	return g.addToken
}

// GetBoard to return the board
func (g *Game) GetBoard() *board.Board {
	return g.board
}

// GetPlayers to return the players
func (g *Game) GetPlayers() []*player.Player {
	return g.players
}

// GetAnalyzer ti return the analyzer
func (g *Game) GetAnalyzer() *analyzer.Analyzer {
	return g.analyzer
}

// GetGameStatus to return the current game status
func (g *Game) GetGameStatus() string {
	return g.gameStatus
}

// GetCurrentPlayer to return the current player
func (g *Game) GetCurrentPlayer() *player.Player {
	return g.currentPlayer
}

// Play to check if the token is added or not, if any player won or if the board is full
func (g *Game) Play(boardNumber int) {
	var tempPlayer *player.Player

	g.addToken = g.addXO(boardNumber, g.currentPlayer.GetMark())
	if g.addToken == false {
		return
	}

	if g.GetAnalyzer().CheckStatus(g.currentPlayer.GetMark()) {
		g.gameStatus = gamestatus.Win
		return
	}
	if g.GetBoard().CheckIfBoardFull() {
		g.gameStatus = gamestatus.Draw
		return
	}
	tempPlayer = g.currentPlayer
	g.currentPlayer = g.nextPlayer
	g.nextPlayer = tempPlayer
}

//addXO to check if the mark can be added to the cell or not
func (g *Game) addXO(boardNumber int, mark string) bool {

	cellCounter := 0

	for i := 0; i < g.GetBoard().GetSize(); i++ {
		for j := 0; j < g.GetBoard().GetSize(); j++ {
			cellCounter++
			if cellCounter == boardNumber {
				if g.GetBoard().GetCells()[i][j].CheckCellTaken() {
					return false
				}
				g.GetBoard().GetCells()[i][j].SetMark(mark)
				return true
			}
		}
	}
	return false
}
