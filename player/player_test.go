package player

import (
	"testing"

	"github.com/sejalnaik/tic-tac-toe/mark"
)

func TestGetName(t *testing.T) {
	player := NewPlayer()
	actual := player.GetName()
	expected := "sejal"

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}

func TestSetMark(t *testing.T) {
	player := NewPlayer()
	player.SetMark(mark.O)
	actual := player.GetMark()
	expected := "O"

	if actual != expected {
		t.Error("Actual is :", actual, " but expected is :", expected)
	}
}
