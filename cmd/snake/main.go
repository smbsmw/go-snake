package main

import (
	"snake"
	"snake/control"
)

func main() {
	controller := control.NewKeyboardController()
	g := snake.NewGame(controller)
	g.Run()
}
