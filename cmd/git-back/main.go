package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/yagince/gitx"
	"regexp"
	"time"
)

type Context struct {
	git                *gitx.Git
	logs               []*gitx.Reflog
	filteredLogs       []*gitx.Reflog
	selectedLineNumber int
	query              string
}

func NewContext(git *gitx.Git, logs []*gitx.Reflog) *Context {
	return &Context{
		git:          git,
		logs:         logs,
		filteredLogs: logs,
	}
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
	if c.selectedLineNumber < len(c.filteredLogs)-1 {
		c.selectedLineNumber += 1
	} else {
		c.selectedLineNumber = 0
	}
	return c.selectedLineNumber
}

func (c *Context) AddQuery(char rune) string {
	if char != 0 {
		c.query += string(char)
	}
	return c.query
}

func (c *Context) EraseQuery(num int) string {
	queryLength := len(c.query)

	if queryLength < num {
		num = queryLength
	}

	c.query = c.query[0 : queryLength-num]
	return c.query
}

func (c Context) QueryRegexp() (*regexp.Regexp, error) {
	return regexp.Compile(c.query)
}

func (c *Context) FilteredLogs() []*gitx.Reflog {
	regexp, err := c.QueryRegexp()
	if err != nil {
		return c.filteredLogs
	}

	logs := make([]*gitx.Reflog, 0)
	for _, log := range c.logs {
		if regexp.MatchString(log.String()) {
			logs = append(logs, log)
		}
	}
	c.filteredLogs = logs

	if len(logs)-1 < c.selectedLineNumber {
		c.selectedLineNumber = len(logs) - 1
	}
	return logs
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
	switch key {
	case termbox.KeyCtrlN, termbox.KeyArrowDown:
		c.Down()
	case termbox.KeyCtrlP, termbox.KeyArrowUp:
		c.Up()
	case termbox.KeyCtrlH, termbox.KeyDelete:
		c.EraseQuery(1)
	}

	clear()

	var y int
	queryPrefix := "QUERY >> "
	drawLine(0, y, queryPrefix, termbox.ColorMagenta)
	drawLine(len(queryPrefix), y, c.query, termbox.ColorDefault)
	y += 1
	drawLine(0, y, "Press ESC or Ctrl+C to exit.", termbox.ColorDefault)
	y += 1
	drawLine(0, y, fmt.Sprintf("-- %d logs", len(c.logs)), termbox.ColorDefault)

	drawLogs(1, y, c)

	flush()
}

func drawLogs(x, y int, c *Context) {
	for i, log := range c.FilteredLogs() {
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

func pollEvent(context *Context, ch chan<- bool) {
	draw(context)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				ch <- true
				return
			case termbox.KeyEnter:
				// TODO: checkout reflog revision
				ch <- true
				return
			default:
				context.AddQuery(ev.Ch)
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

	termbox.SetInputMode(termbox.InputAlt)

	git := gitx.NewGit("./")
	logs := git.Reflog(0)
	context := NewContext(git, logs)
	ch := make(chan bool)

	go pollEvent(context, ch)
	<-ch
}
