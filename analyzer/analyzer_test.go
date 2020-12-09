package analyzer

import (
	"testing"
	"tic-tac-toe/board"
)

func TestCheckRowsMatch(t *testing.T) {
	var board = board.NewBoard(3)
	var analyzer = NewAnalyzer(board)
	analyzer.SetMark("O")

	board.GetCells()[0][0].SetMark("O")
	board.GetCells()[1][0].SetMark("O")
	board.GetCells()[2][0].SetMark("O")
	actual := analyzer.checkRowsMatch()

	expected := false

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckColumnsMatch(t *testing.T) {
	var board = board.NewBoard(3)
	var analyzer = NewAnalyzer(board)
	analyzer.SetMark("O")

	board.GetCells()[0][0].SetMark("O")
	board.GetCells()[0][1].SetMark("O")
	board.GetCells()[0][2].SetMark("O")
	actual := analyzer.checkColumnsMatch()

	expected := false

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckDiagonalMatch(t *testing.T) {
	var board = board.NewBoard(3)
	var analyzer = NewAnalyzer(board)
	analyzer.SetMark("O")

	board.GetCells()[0][0].SetMark("O")
	board.GetCells()[1][0].SetMark("O")
	board.GetCells()[2][0].SetMark("O")
	actual := analyzer.checkDiagonalMatch()

	expected := false

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckReverseDiagonalMatch(t *testing.T) {
	var board = board.NewBoard(3)
	var analyzer = NewAnalyzer(board)
	analyzer.SetMark("O")

	board.GetCells()[0][0].SetMark("O")
	board.GetCells()[1][1].SetMark("O")
	board.GetCells()[2][2].SetMark("O")
	actual := analyzer.checkReverseDiagonalMatch()

	expected := false

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckStatus(t *testing.T) {
	var board = board.NewBoard(3)
	var analyzer = NewAnalyzer(board)

	board.GetCells()[0][0].SetMark("X")
	board.GetCells()[0][2].SetMark("O")
	board.GetCells()[2][0].SetMark("X")
	board.GetCells()[1][0].SetMark("O")
	board.GetCells()[2][2].SetMark("X")
	board.GetCells()[2][1].SetMark("O")
	board.GetCells()[1][1].SetMark("X")
	actual := analyzer.CheckStatus("X")

	expected := true

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}
