package snake

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

const (
	wallBlock rune = ' '
	snakeBodyBlock
	foodBlock rune = '🍒'
)

type render struct {
	w *world
}

func init() {
	err := termbox.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

func newRender(w *world) *render {
	return &render{w}
}

func (r *render) clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (r *render) flush() {
	termbox.Flush()
}

func (r *render) draw() {
	r.clear()

	r.drawInfo()
	r.drawWall()
	r.drawFood()
	r.drawSnake()

	r.flush()
}

func (r *render) drawSnake() {
	for _, b := range r.w.snake.body {
		termbox.SetCell(b.X, b.Y, snakeBodyBlock, termbox.ColorGreen, termbox.ColorGreen)
	}
}

func (r *render) drawText(x, y int, text string) {
	for i, r := range text {
		termbox.SetChar(x+i, y, r)
	}
}

func (r *render) drawInfo() {
	r.drawText(r.w.wall.width+10, 0, "SNAKE GAME")
	r.drawText(r.w.wall.width+10, 1, fmt.Sprintf("SCORE: %d", r.w.score()))
	switch r.w.status {
	case statusPaused:
		r.drawPopup("PAUSED")
	case statusGameover:
		r.drawPopup("GAME OVER")
	}
}

func (r *render) drawPopup(text string) {
	x, y := r.w.wall.width/2-len(text)/2, r.w.wall.height/2
	if x%2 != 0 {
		x--
	}
	r.drawText(x, y, text)
}

func (r *render) drawWall() {
	for i := 0; i <= r.w.wall.width; i++ {
		termbox.SetCell(i, 0, wallBlock, termbox.ColorDarkGray, termbox.ColorDarkGray)
		termbox.SetCell(i, r.w.wall.height, wallBlock, termbox.ColorDarkGray, termbox.ColorDarkGray)
	}

	for i := 1; i < r.w.wall.height; i++ {
		termbox.SetCell(0, i, wallBlock, termbox.ColorDarkGray, termbox.ColorDarkGray)
		termbox.SetCell(r.w.wall.width, i, wallBlock, termbox.ColorDarkGray, termbox.ColorDarkGray)
	}
}

func (r *render) drawFood() {
	termbox.SetCell(r.w.food.coord.X, r.w.food.coord.Y, foodBlock, termbox.ColorDefault, termbox.ColorDefault)
}
