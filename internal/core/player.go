package core

type Player struct {
	Name string
	Alive bool
	Hand   CardStack
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Hand: CardStack{Cards: make([]Card, 0)},
	}
}
