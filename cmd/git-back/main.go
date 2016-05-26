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

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func flush() {
	termbox.Flush()
}

func draw(b []*gitx.Reflog) {
	drawWithKey(0, b)
}

func drawWithKey(key termbox.Key, logs []*gitx.Reflog) {
	clear()
	var y int
	drawLine(0, y, "Press ESC or Ctrl+C to exit.", termbox.ColorDefault)
	y += 1
	drawLine(0, y, fmt.Sprintf("-- %d logs", len(logs)), termbox.ColorDefault)

	switch key {
	case termbox.KeyCtrlN, termbox.KeyArrowDown:
		// TODO: 選択列下げる
		drawLogs(1, y, logs)
	case termbox.KeyCtrlP, termbox.KeyArrowUp:
		// TODO: 選択列上げる
		drawLogs(1, y, logs)
	default:
		drawLogs(1, y, logs)
	}

	flush()
}

func drawLogs(x, y int, logs []*gitx.Reflog) {
	for _, log := range logs {
		y += 1

		var color termbox.Attribute
		// if i == b.Selected {
		// 	color = termbox.ColorGreen
		// } else if i == b.Current {
		// 	color = termbox.ColorMagenta
		// } else {
		color = termbox.ColorDefault
		// }

		drawLine(1, y, log.String(), color)
	}
}

func pollEvent(git *gitx.Git) {
	logs := git.Reflog(0)
	draw(logs)
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
				// 	return
				// }
			default:
				drawWithKey(ev.Key, logs)
			}
		default:
			draw(logs)
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
