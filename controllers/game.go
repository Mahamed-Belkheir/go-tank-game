package controllers

import "go-tanks-game/models"

type Game struct {
	player1 *models.Player
	player2 *models.Player

	GameMap *models.GameMap
}

func NewGame(player1 string, player2 string) *Game {
	limit := 5

	newGame := Game{
		player1: models.NewPlayer(player1, 1, limit),
		player2: models.NewPlayer(player2, 2, limit),
		GameMap: models.NewGameMap(),
	}
	for i := 0; i < limit; i++ {
		tank := newGame.player1.Tanks[i]
		newGame.GameMap.Matrix[0][i] = tank
		tank.GameMap = newGame.GameMap
		tank.X, tank.Y = i, 0
		tank = newGame.player2.Tanks[i]
		newGame.GameMap.Matrix[limit-1][i] = tank
		tank.GameMap = newGame.GameMap
		tank.X, tank.Y = i, limit-1
	}
	return &newGame
}
