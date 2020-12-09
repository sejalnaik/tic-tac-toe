//Package cell is for creating and managing cell
package cell

import (
	"tic-tac-toe/mark"
)

// Cell is defined
type Cell struct {
	mark string
}

//NewCell is for craeting new cell
func NewCell() *Cell {
	var cellTest = &Cell{
		mark: mark.BLANK,
	}
	return cellTest
}

//GetMark is for returning value of mark
func (c *Cell) GetMark() string {
	return c.mark
}

//SetMark is setting mark of cell
func (c *Cell) SetMark(mark string) {
	c.mark = mark
}

//CheckCellTaken is for checking whether the mark is blank or not
func (c *Cell) CheckCellTaken() bool {
	if c.mark != mark.BLANK {
		return true
	}
	return false
}
