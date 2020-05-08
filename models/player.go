package models

type Player struct {
	Name  string
	Id    int
	Tanks map[int]*Tank
}

func NewPlayer(name string, id int, tankNum int) *Player {
	newPlayer := Player{
		Name:  name,
		Id:    id,
		Tanks: make(map[int]*Tank),
	}
	for i := 0; i <= tankNum; i++ {
		newPlayer.Tanks[i] = NewTank(&newPlayer, i)
	}
	return &newPlayer
}
