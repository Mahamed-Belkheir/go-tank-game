package main

import (
	"go-tanks-game/controllers"
)

func main() {
	game := controllers.NewGame("bob", "john")
	game.GameMap.View()
	tank := game.GameMap.Matrix[4][0]
	tank.Move(1)
	tank.Move(2)
	tank.Move(2)
	tank.Move(2)
	tank.Move(1)
	tank.Move(1)
	tank.Shoot(1)
	tank.Shoot(1)
	tank.Move(1)
	tank.Shoot(4)
	tank.Shoot(4)

	game.GameMap.View()
}
