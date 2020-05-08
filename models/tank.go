package models

import (
	"fmt"
)

type Tank struct {
	X       int
	Y       int
	Owner   *Player
	GameMap *GameMap
	Id      int

	Health int
	Attack int
}

var directionsOffset = map[int][2]int{
	1: [2]int{0, -1},
	2: [2]int{1, 0},
	3: [2]int{0, 1},
	4: [2]int{-1, 0},
}

func (t *Tank) Shoot(direction int) bool {
	offset := directionsOffset[direction]
	x := t.X + offset[0]
	y := t.Y + offset[1]
	limit := len(t.GameMap.Matrix)
	fmt.Println(x, y)
	if x < 0 || x > limit-1 || y < 0 || y > limit-1 || t.GameMap.Matrix[y][x] == nil {
		return false
	}
	target := t.GameMap.Matrix[y][x]
	target.Defend(t.Attack)
	return true
}

func (t *Tank) Defend(damage int) bool {
	t.Health -= damage
	if t.Health > 0 {
		return false
	}
	t.GameMap.Matrix[t.Y][t.X] = nil
	delete(t.Owner.Tanks, t.Id)
	return true
}

func (t *Tank) Move(direction int) bool {
	offset := directionsOffset[direction]
	x := t.X + offset[0]
	y := t.Y + offset[1]
	limit := len(t.GameMap.Matrix)
	if x < 0 || x > limit-1 || y < 0 || y > limit-1 || t.GameMap.Matrix[y][x] != nil {
		return false
	}
	t.GameMap.Matrix[t.Y][t.X] = nil
	t.GameMap.Matrix[y][x] = t
	t.X = x
	t.Y = y
	return true

}

func NewTank(player *Player, id int) *Tank {
	tank := Tank{
		Owner:  player,
		Id:     id,
		Health: 100,
		Attack: 50,
	}
	return &tank
}
