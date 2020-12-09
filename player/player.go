package player

// Player to create type for player
type Player struct {
	playerName string
	playerMark string
	playerID   int
}

// NewPlayer to craete a struct of type Player
func NewPlayer() *Player {
	return &Player{
		playerName: "",
		playerMark: "*",
		playerID:   0,
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

// SetID to set the id of the player
func (p *Player) SetID(id int) {
	p.playerID = id
}

// GetID to set the id of the player
func (p *Player) GetID() int {
	return p.playerID
}

// GetName to return the name of the player
func (p *Player) GetName() string {
	return p.playerName
}

// SetName to set the name of the player
func (p *Player) SetName(name string) {
	p.playerName = name
}
