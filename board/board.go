package board

import (
	"github.com/sejalnaik/tic-tac-toe/cell"
	"github.com/sejalnaik/tic-tac-toe/mark"
)

// Board struct
type Board struct {
	cells [][]*cell.Cell
	size  int
}

//NewBoard to create new board and initialize it with a provided size
func NewBoard(size int) *Board {

	tempCells := make([][]*cell.Cell, size)

	for i := range tempCells {
		tempCells[i] = make([]*cell.Cell, size)
		for j := range tempCells[i] {
			tempCell := cell.NewCell()
			tempCells[i][j] = tempCell
		}
	}

	var boardTest = &Board{
		cells: tempCells,
		size:  size,
	}
	return boardTest
}

// GetSize to get the size of board
func (b *Board) GetSize() int {
	return b.size
}

// GetCells to retuen the cells of board
func (b *Board) GetCells() [][]*cell.Cell {
	return b.cells
}

// CheckIfBoardFull to chcek if board is full(all cell marked with X and O)
func (b *Board) CheckIfBoardFull() bool {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.cells[i][j].GetMark() == mark.Blank {
				return false
			}
		}
	}
	return true
}
