package snake

import (
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

// game represents model of game
type game struct {
	world  *world
	render *render
}

// NewGame creates new instance of game
func NewGame(c Controller) *game {
	w := &world{
		controller: c,
		cmd:        make(chan CommandID),
		status:     statusActive,
		food:       &food{coord{10, 10}},
		wall:       &wall{50, 15},
		snake: &snake{
			direction: coord{1, 0},
			body: []coord{
				{3, 3},
				{4, 3},
				{5, 3},
			},
		},
		ticker: time.NewTicker(150 * time.Millisecond),
	}

	return &game{
		world:  w,
		render: newRender(w),
	}
}

type statusID int

const (
	statusActive statusID = iota
	statusPaused
	statusGameover
)

// Run runs the game
func (g *game) Run() {
	go g.world.controller.Control(g.world.cmd)

	for {
		select {
		case d := <-g.world.cmd:
			switch {
			case d == UP:
				g.world.snake.direction = coord{0, -1}
			case d == DOWN:
				g.world.snake.direction = coord{0, 1}
			case d == LEFT:
				g.world.snake.direction = coord{-1, 0}
			case d == RIGHT:
				g.world.snake.direction = coord{1, 0}
			case d == ESC:
				if g.world.status == statusPaused {
					g.world.status = statusActive
				} else {
					g.world.status = statusPaused
				}
			case d == QUIT:
				g.Exit()

			}
		case <-g.world.ticker.C:
			switch g.world.status {
			case statusActive:
				g.world.onTick()
			case statusPaused:

			}
			g.render.draw()
		}
	}

}

func (g *game) Exit() {
	termbox.Close()
	os.Exit(0)
}