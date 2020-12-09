package cell

import (
	"testing"

	"github.com/sejalnaik/tic-tac-toe/mark"
)

func TestNewCell(t *testing.T) {
	var cellOne = NewCell()
	actual := cellOne.mark
	var expected = "*"

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestGetMark(t *testing.T) {
	var cellOne = NewCell()
	actual := cellOne.GetMark()
	expected := "*"
	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestSetMark(t *testing.T) {
	var cellOne = NewCell()
	cellOne.SetMark(mark.X)
	actual := cellOne.GetMark()
	expected := "X"
	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestCheckCellTaken(t *testing.T) {
	var cellOne = NewCell()
	cellOne.SetMark(mark.X)
	actual := cellOne.CheckCellTaken()
	expected := true

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}
