package gamestart

import (
	"fmt"
	"log"

	"github.com/sejalnaik/tic-tac-toe/game"
	gamestatus "github.com/sejalnaik/tic-tac-toe/gameStatus"
	"github.com/sejalnaik/tic-tac-toe/mark"
	"github.com/sejalnaik/tic-tac-toe/player"
)

// Start to start the game
func Start() {
	// take size input form user
	fmt.Println("Enter size of the game")
	var size int
	_, errorSize := fmt.Scan(&size)
	if errorSize != nil {
		log.Fatal("Error occurred:", errorSize)
	}

	// take players info form users
	players := make([]*player.Player, 2)
	fmt.Println("Enter player 1 name")
	var name1 string
	_, errorName1 := fmt.Scan(&name1)
	if errorName1 != nil {
		log.Fatal("Error occurred:", errorName1)
	}
	player1 := player.NewPlayer(name1)
	player1.SetMark(mark.X)

	fmt.Println("Enter player 2 name")
	var name2 string
	_, errorName2 := fmt.Scan(&name2)
	if errorName2 != nil {
		log.Fatal("Error occurred:", errorName2)
	}
	player2 := player.NewPlayer(name2)
	player2.SetMark(mark.O)

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
			fmt.Println(myGame.GetCurrentPlayer().GetName(), "play", myGame.GetCurrentPlayer().GetMark())
			fmt.Println("Enter the cell number to be marked")
			var userCellInput int
			_, errorUserCellInput := fmt.Scan(&userCellInput)
			if errorUserCellInput != nil {
				log.Fatal("Error occurred:", errorUserCellInput)
			}
			if userCellInput > size*size || userCellInput < 1 {
				fmt.Println("Please enter a number between 1 and ", size*size)
				continue
			}
			myGame.Play(userCellInput)
		}
		displayGame(myGame)
		if myGame.GetGameStatus() == gamestatus.WIN {
			fmt.Println(myGame.GetCurrentPlayer().GetName(), "WINS!!!")
			return
		}
		if myGame.GetGameStatus() == gamestatus.DRAW {
			fmt.Println("Noone wins, its a draw")
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
