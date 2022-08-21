package control

import (
	"github.com/nsf/termbox-go"
	"log"

	"snake"
)


type keyboard struct{}

func NewKeyboardController() *keyboard {
	return &keyboard{}
}

func (k *keyboard) Control(ids chan<- snake.CommandID) {
	defer termbox.Close()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				ids <- snake.UP
			case termbox.KeyArrowDown:
				ids <- snake.DOWN
			case termbox.KeyArrowLeft:
				ids <- snake.LEFT
			case termbox.KeyArrowRight:
				ids <- snake.RIGHT
			case termbox.KeyEsc:
				ids <- snake.ESC
			case termbox.KeyCtrlC:
				ids <- snake.QUIT
			default:
			}
		case termbox.EventError:
			log.Fatalln(ev.Err)
		}
	}
}
