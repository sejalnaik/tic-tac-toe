package board

import (
	"fmt"
	"testing"

	"github.com/sejalnaik/tic-tac-toe/mark"
)

func TestNewBoard(t *testing.T) {
	var boardOne = NewBoard(3)
	for i := range boardOne.GetCells() {
		for j := range boardOne.GetCells()[i] {
			fmt.Print(boardOne.GetCells()[i][j])
		}
		fmt.Println()
	}
}

func TestGetSize(t *testing.T) {
	var boardOne = NewBoard(3)
	actual := boardOne.GetSize()
	var expected = 3

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckIfBoardFull(t *testing.T) {
	var boardOne = NewBoard(3)
	for i := range boardOne.GetCells() {
		for j := range boardOne.GetCells()[i] {
			boardOne.GetCells()[i][j].SetMark(mark.O)
		}
		fmt.Println()
	}
	actual := boardOne.CheckIfBoardFull()
	var expected = true

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}
