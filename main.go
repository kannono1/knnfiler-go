package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

var colors = []termbox.Attribute{
	termbox.ColorDefault,
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
}

func drawX(x, y int, str string, fgColor int, bgColor int) {
	runes := []rune(str)
	for _, r := range runes {
		termbox.SetCell(x, y, r, colors[fgColor], colors[bgColor])
		x += runewidth.RuneWidth(r)
	}
}

func redraw() {
	w, _ := termbox.Size()
	w2x := int(w / 2)
	termbox.Clear(coldef, coldef)
	drawX(0, 0, "Hello", 0, 0)
	drawX(0, 1, "         ", 0, 1)
	drawX(0, 2, "こんにちは", 1, 3)
	drawX(0, 3, "abcあいう", 0, 0)
	drawX(0, 4, "a c い | う", 0, 0)
	drawX(w2x, 2, "こんにちは", 0, 0)
	drawX(w2x, 3, "abcあいう", 0, 0)
	drawX(w2x, 4, "a c い | う", 0, 0)
	termbox.Flush()
}

func main() {
	print("Hello knnfiler")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	redraw()
MAINLOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			}
		}
		redraw()
	}
}
