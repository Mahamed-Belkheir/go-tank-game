package models

import "fmt"

type GameMap struct {
	Matrix [5][5]*Tank
}

func NewGameMap() *GameMap {
	return &GameMap{}
}

func (g *GameMap) View() {
	fmt.Println("viewing map")
	for _, row := range g.Matrix {
		for _, tank := range row {
			if tank != nil {
				fmt.Print("[", tank.Owner.Id, "]")
			} else {
				fmt.Print("[ ]")
			}
		}
		fmt.Println("")
	}
}
