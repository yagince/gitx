package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os/exec"
	"strings"
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
		} else if i == b.current {
			color = termbox.ColorMagenta
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
	drawLine(0, y, "Press ESC or Ctrl+C to exit.", termbox.ColorDefault)
	y += 1
	drawLine(0, y, fmt.Sprintf("-- %d branches", len(b.values)), termbox.ColorDefault)

	switch key {
	case termbox.KeyCtrlN:
		b.Down()
		drawBranches(1, y, b)
	case termbox.KeyCtrlP:
		b.Up()
		drawBranches(1, y, b)
	default:
		drawBranches(1, y, b)
	}

	termbox.Flush()
}

func pollEvent(branches *Branches) {
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

	git := NewGit("./")
	branches := git.Branches()
	pollEvent(branches)
}

type Git struct {
	binary    string
	directory string
}

func NewGit(directory string) *Git {
	var binary string
	var err error
	if binary, err = exec.LookPath("git"); err != nil {
		panic(err)
	}
	return &Git{binary: binary, directory: directory}
}

func (g *Git) Branches() *Branches {
	lf := "\n"
	cmd := exec.Command(g.binary, "branch")

	var out []byte
	var err error
	if out, err = cmd.Output(); err != nil {
		panic(err)
	}

	branches := strings.Split(strings.TrimRight(string(out), lf), lf)
	var current int
	for i, b := range branches {
		if strings.IndexAny(b, "*") == 0 {
			current = i
		}
	}
	return &Branches{values: branches, current: current}
}

type Branches struct {
	values   []string
	current  int
	selected int
}

func (b *Branches) Up() int {
	if b.selected != 0 {
		b.selected -= 1
	}
	return b.selected
}

func (b *Branches) Down() int {
	if (b.selected + 1) < len(b.values) {
		b.selected += 1
	}
	return b.selected
}
