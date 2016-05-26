package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/yagince/gitx"
	"time"
)

type Context struct {
	logs               []*gitx.Reflog
	selectedLineNumber int
}

func (c *Context) Up() int {
	if c.selectedLineNumber > 0 {
		c.selectedLineNumber -= 1
	} else {
		c.selectedLineNumber = len(c.logs) - 1
	}
	return c.selectedLineNumber
}

func (c *Context) Down() int {
	if c.selectedLineNumber < len(c.logs)-1 {
		c.selectedLineNumber += 1
	} else {
		c.selectedLineNumber = 0
	}
	return c.selectedLineNumber
}

func drawLine(x, y int, str string, color termbox.Attribute) {
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func flush() {
	termbox.Flush()
}

func draw(c *Context) {
	drawWithKey(0, c)
}

func drawWithKey(key termbox.Key, c *Context) {
	clear()
	var y int
	drawLine(0, y, "Press ESC or Ctrl+C to exit.", termbox.ColorDefault)
	y += 1
	drawLine(0, y, fmt.Sprintf("-- %d logs", len(c.logs)), termbox.ColorDefault)

	switch key {
	case termbox.KeyCtrlN, termbox.KeyArrowDown:
		c.Down()
		drawLogs(1, y, c)
	case termbox.KeyCtrlP, termbox.KeyArrowUp:
		c.Up()
		drawLogs(1, y, c)
	default:
		drawLogs(1, y, c)
	}

	flush()
}

func drawLogs(x, y int, c *Context) {
	for i, log := range c.logs {
		y += 1

		var color termbox.Attribute
		if i == c.selectedLineNumber {
			color = termbox.ColorGreen
		} else {
			color = termbox.ColorDefault
		}

		drawLine(1, y, log.String(), color)
	}
}

func pollEvent(git *gitx.Git) {
	logs := git.Reflog(0)
	context := &Context{logs: logs}
	draw(context)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			case termbox.KeyEnter:
				// if out, err := git.CheckOut(branches.SelectedBranch()); err != nil {
				// 	clear()
				// 	drawLine(1, 1, string(out), termbox.ColorRed)
				// 	flush()
				// } else {
				return
				// }
			default:
				drawWithKey(ev.Key, context)
			}
		default:
			draw(context)
		}
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			clear()
			drawLine(0, 0, fmt.Sprintf("%v", err), termbox.ColorRed)
			flush()
			time.Sleep(3 * time.Second)
		}
		termbox.Close()
	}()

	git := gitx.NewGit("./")
	pollEvent(git)
}
