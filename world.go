package snake

import (
	"math/rand"
	"time"
)

// world represents playground for game
type world struct {
	ticker     *time.Ticker
	cmd        chan CommandID
	status     statusID
	wall       *wall
	snake      *snake
	food       *food
	controller Controller
}

func (w *world) checkCollision() {
	switch {
	case w.snake.head() == w.food.coord:
		w.snake.eat()
		w.generateFruit()
	case w.wall.onWall(w.snake.head()) || w.snake.headOnBody():
		w.status = statusGameover
		w.ticker.Stop()
	}
}

func (w *world) onTick() {
	w.snake.nextMove()
	w.checkCollision()
}

func (w *world) score() int {
	//Default length is 3
	return w.snake.length() - 3
}

type wall struct {
	width, height int
}

func (w *wall) onWall(c coord) bool {
	return c.X == 0 || c.X == w.width || c.Y == 0 || c.Y == w.height
}

type coord struct {
	X, Y int
}

type food struct {
	coord coord
}

func (w *world) generateFruit() {
	w.food.coord.X = rand.Intn(w.wall.width-1) + 1
	w.food.coord.Y = rand.Intn(w.wall.height-1) + 1
}
