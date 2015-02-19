package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func drawLine(x, y int, str string, color termbox.Attribute) {
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawBranches(x, y int, b *Branches) {
	for i, branch := range b.values {
		y += 1

		var color termbox.Attribute
		if i == b.selected {
			color = termbox.ColorGreen
		} else {
			color = termbox.ColorDefault
		}

		drawLine(1, y, fmt.Sprintf("%0d: %s", i, branch), color)
	}
}

func draw(b *Branches) {
	drawWithKey(0, b)
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func drawWithKey(key termbox.Key, b *Branches) {
	clear()
	var y int
	drawLine(0, y, fmt.Sprintf("Press ESC or Ctrl+C to exit. %d %s", key, time.Now()), termbox.ColorDefault)

	switch key {
	case termbox.KeyCtrlN:
		b.down()
		drawBranches(1, y, b)
	case termbox.KeyCtrlP:
		b.up()
		drawBranches(1, y, b)
	default:
		drawBranches(1, y, b)
	}

	termbox.Flush()
}

func pollEvent() {
	branches := &Branches{values: []string{"hoge", "foo", "bar", "buzz"}}
	draw(branches)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			case termbox.KeyEnter:
				// TODO: Change Branch
				return
			default:
				drawWithKey(ev.Key, branches)
			}
		default:
			draw(branches)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	pollEvent()
}

type Git struct {
	directory string
}

func (g *Git) branches() {
}

type Branches struct {
	values   []string
	current  int
	selected int
}

func (b *Branches) up() int {
	if b.selected != 0 {
		b.selected -= 1
	}
	return b.selected
}

func (b *Branches) down() int {
	if (b.selected + 1) < len(b.values) {
		b.selected += 1
	}
	return b.selected
}
