package gamestart

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/sejalnaik/tic-tac-toe/game"
	gamestatus "github.com/sejalnaik/tic-tac-toe/gameStatus"
	"github.com/sejalnaik/tic-tac-toe/mark"
	"github.com/sejalnaik/tic-tac-toe/player"
)

// Start to start the game
func Start() {
	// take size input form user
	var sizeStr string = "8"
	var size int
	fmt.Println("Enter size of the board")
	for okSize := true; okSize; okSize = (okSize != isSizeNumber(sizeStr)) {
		if isSizeNumber(sizeStr) == false {
			fmt.Print(("Please enter only numbers"))
		}
	}

	size, _ = strconv.Atoi(sizeStr)

	players := make([]*player.Player, 2)

	// take player1 info form user
	player1 := player.NewPlayer()
	player1.SetMark(mark.X)
	player1.SetID(1)
	setPlayer(player1)

	// take player2 info form user
	player2 := player.NewPlayer()
	player2.SetMark(mark.O)
	player2.SetID(2)
	setPlayer(player2)

	// put players in a slice
	players[0] = player1
	players[1] = player2

	//create new game
	myGame := game.NewGame(players, size)

	//game starts here
	fmt.Println("********************Game starts*****************************")
	displayGame(myGame)
	for ok1 := true; ok1; ok1 = (myGame.GetGameStatus() == gamestatus.INPROGRESS) {
		for ok2 := true; ok2; ok2 = (myGame.IsAddToken() == false) {
			if myGame.IsAddToken() == false {
				fmt.Println("Cell taken, enter some other board number")
				displayGame(myGame)
			}
			fmt.Println("Player", myGame.GetCurrentPlayer().GetID(), ":", myGame.GetCurrentPlayer().GetName(), "play", myGame.GetCurrentPlayer().GetMark())
			fmt.Println("Enter the cell number to be marked")
			var userCellInputStr string

			_, errorUserCellInput := fmt.Scan(&userCellInputStr)
			if errorUserCellInput != nil {
				fmt.Println("Error occurred:", errorUserCellInput)
			}
			userCellInputInt, err := strconv.Atoi(userCellInputStr)
			if err != nil {
				fmt.Println("Please enter a number between 1 and ", size*size)
				continue
			}

			if userCellInputInt > size*size || userCellInputInt < 1 {
				fmt.Println("Please enter a number between 1 and ", size*size)
				continue
			}
			myGame.Play(userCellInputInt)
		}
		displayGame(myGame)
		if myGame.GetGameStatus() == gamestatus.WIN {
			fmt.Println(myGame.GetCurrentPlayer().GetName(), "WINS!!!")
			return
		}
		if myGame.GetGameStatus() == gamestatus.DRAW {
			fmt.Println("No one wins, its a draw")
			return
		}
	}
}

//display the game board
func displayGame(tempGame *game.Game) {
	for i := 0; i < tempGame.GetBoard().GetSize(); i++ {
		for j := 0; j < tempGame.GetBoard().GetSize(); j++ {
			fmt.Print(tempGame.GetBoard().GetCells()[i][j].GetMark(), " ")
		}
		fmt.Println()
	}
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func setPlayer(player *player.Player) {
	var name = "a"
	for okName := true; okName; okName = (okName != isLetter(name)) {
		if isLetter(name) == false {
			fmt.Println("Please enter only alphabets")
		}
		fmt.Println("Enter player", player.GetID(), "name")
		_, errorName1 := fmt.Scan(&name)
		if errorName1 != nil {
			fmt.Print("Error occurred:", errorName1)
		}
		player.SetName(name)
	}
}

func isSizeNumber(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
