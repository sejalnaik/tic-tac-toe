package player

// Player to create type for player
type Player struct {
	playerName string
	playerMark string
}

// NewPlayer to craete a struct of type Player
func NewPlayer(name string) *Player {
	return &Player{
		playerName: name,
		playerMark: "*",
	}
}

// GetMark to return the mark of the player
func (p *Player) GetMark() string {
	return p.playerMark
}

// SetMark to set the mark of the player
func (p *Player) SetMark(mark string) {
	p.playerMark = mark
}

// GetName to return the name of the player
func (p *Player) GetName() string {
	return p.playerName
}
