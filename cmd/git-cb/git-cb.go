package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/yagince/gitx"
	"time"
)

func drawLine(x, y int, str string, color termbox.Attribute) {
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawBranches(x, y int, b *gitx.Branches) {
	for i, branch := range b.Values {
		y += 1

		var color termbox.Attribute
		if i == b.Selected {
			color = termbox.ColorGreen
		} else if i == b.Current {
			color = termbox.ColorMagenta
		} else {
			color = termbox.ColorDefault
		}

		drawLine(1, y, fmt.Sprintf("%0d: %s", i, branch), color)
	}
}

func draw(b *gitx.Branches) {
	drawWithKey(0, b)
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func flush() {
	termbox.Flush()
}

func drawWithKey(key termbox.Key, b *gitx.Branches) {
	clear()
	var y int
	drawLine(0, y, "Press ESC or Ctrl+C to exit.", termbox.ColorDefault)
	y += 1
	drawLine(0, y, fmt.Sprintf("-- %d branches", len(b.Values)), termbox.ColorDefault)

	switch key {
	case termbox.KeyCtrlN, termbox.KeyArrowDown:
		b.Down()
		drawBranches(1, y, b)
	case termbox.KeyCtrlP, termbox.KeyArrowUp:
		b.Up()
		drawBranches(1, y, b)
	default:
		drawBranches(1, y, b)
	}

	flush()
}

func pollEvent(git *gitx.Git) {
	branches := git.Branches()
	draw(branches)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			case termbox.KeyEnter:
				if out, err := git.CheckOut(branches.SelectedBranch()); err != nil {
					clear()
					drawLine(1, 1, string(out), termbox.ColorRed)
					flush()
				} else {
					return
				}
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

	defer func() {
		if r := recover(); r != nil {
			clear()
			drawLine(0, 0, fmt.Sprintf("%v", r), termbox.ColorRed)
			flush()
			time.Sleep(3 * time.Second)
		}
		termbox.Close()
	}()

	git := gitx.NewGit("./")
	pollEvent(git)
}
