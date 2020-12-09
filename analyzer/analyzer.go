package analyzer

import (
	"github.com/sejalnaik/tic-tac-toe/board"
)

// Analyzer to create anylyzer structs
type Analyzer struct {
	board     *board.Board
	countMark int
	markCell  string
}

// NewAnalyzer to crate a new instance of analyzer type
func NewAnalyzer(b *board.Board) *Analyzer {
	var tempBoard = b
	var tempCountMark = 0
	var tempMarkCell = "*"
	var tempAnalyzer = &Analyzer{
		board:     tempBoard,
		countMark: tempCountMark,
		markCell:  tempMarkCell,
	}
	return tempAnalyzer
}

// GetBoard to return the board
func (a *Analyzer) GetBoard() *board.Board {
	return a.board
}

// SetMark to set the mark of analyzer
func (a *Analyzer) SetMark(mark string) {
	a.markCell = mark
}

// checkRowsMatch to check matching rows
func (a *Analyzer) checkRowsMatch() bool {
	a.countMark = 0
	for i := 0; i < a.GetBoard().GetSize(); i++ {
		for j := 0; j < a.GetBoard().GetSize(); j++ {
			if a.GetBoard().GetCells()[i][j].GetMark() == a.markCell {
				a.countMark++
			}
		}
		if a.countMark == a.GetBoard().GetSize() {
			return true
		} else {
			a.countMark = 0
		}
	}
	return false
}

// checkColumnsMatch to check matching columns
func (a *Analyzer) checkColumnsMatch() bool {
	a.countMark = 0
	for i := 0; i < a.GetBoard().GetSize(); i++ {
		for j := 0; j < a.GetBoard().GetSize(); j++ {
			if a.GetBoard().GetCells()[j][i].GetMark() == a.markCell {
				a.countMark++
			}
		}
		if a.countMark == a.GetBoard().GetSize() {
			return true
		} else {
			a.countMark = 0
		}
	}
	return false
}

// checkDiagonalMatch to check forward diagonal
func (a *Analyzer) checkDiagonalMatch() bool {
	a.countMark = 0
	for i := 0; i < a.GetBoard().GetSize(); i++ {
		if a.GetBoard().GetCells()[i][i].GetMark() == a.markCell {
			a.countMark++
		}
	}
	if a.countMark == a.GetBoard().GetSize() {
		return true
	} else {
		a.countMark = 0
	}
	return false
}

// checkReverseDiagonalMatch to check reverse diagonals
func (a *Analyzer) checkReverseDiagonalMatch() bool {
	a.countMark = 0
	for i, j := 0, a.GetBoard().GetSize()-1; i < a.GetBoard().GetSize() && j >= 0; i, j = i+1, j-1 {
		if a.GetBoard().GetCells()[i][j].GetMark() == a.markCell {
			a.countMark++
		}
	}
	if a.countMark == a.GetBoard().GetSize() {
		return true
	} else {
		a.countMark = 0
	}
	return false
}

// CheckStatus to check any matches
func (a *Analyzer) CheckStatus(mark string) bool {
	a.markCell = mark
	if a.checkRowsMatch() {
		return true
	}
	if a.checkColumnsMatch() {
		return true
	}
	if a.checkDiagonalMatch() {
		return true
	}
	if a.checkReverseDiagonalMatch() {
		return true
	}
	return false
}
