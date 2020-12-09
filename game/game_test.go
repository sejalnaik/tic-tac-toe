package game

import (
	"fmt"
	"testing"
	"tic-tac-toe/mark"
	"tic-tac-toe/player"
)

func TestAddXO(t *testing.T) {

	player1 := player.NewPlayer("sejal")
	player2 := player.NewPlayer("rachel")

	players := make([]*player.Player, 2)

	players[0] = player1
	players[1] = player2

	var game = NewGame(players, 3)

	game.GetBoard().GetCells()[0][0].SetMark("X")
	game.GetBoard().GetCells()[0][2].SetMark("O")
	game.GetBoard().GetCells()[2][0].SetMark("X")
	game.GetBoard().GetCells()[1][0].SetMark("O")
	game.GetBoard().GetCells()[2][2].SetMark("X")
	game.GetBoard().GetCells()[2][1].SetMark("O")
	game.GetBoard().GetCells()[1][1].SetMark("X")

	actual := game.addXO(1, "X")

	fmt.Println(game.GetBoard().GetCells()[0][0].GetMark())

	expected := false

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestPlay(t *testing.T) {
	player1 := player.NewPlayer("sejal")
	player2 := player.NewPlayer("rachel")

	players := make([]*player.Player, 2)

	players[0] = player1
	player1.SetMark(mark.X)
	players[1] = player2
	player2.SetMark(mark.O)

	var game = NewGame(players, 3)

	game.Play(1)
	disPlay(game)
	game.Play(2)
	disPlay(game)
	game.Play(2)
	disPlay(game)
	game.Play(3)
	disPlay(game)
	/*game.Play(1)
	disPlay(game)*/

}

func disPlay(game *Game) {
	for i := 0; i < game.GetBoard().GetSize(); i++ {
		for j := 0; j < game.GetBoard().GetSize(); j++ {
			fmt.Print(game.GetBoard().GetCells()[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
