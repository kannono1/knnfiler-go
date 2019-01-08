package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

const coldef = termbox.ColorDefault

func drawBox(x, y int) {
	termbox.Clear(coldef, coldef)
	termbox.SetCell(x, y, '┏', coldef, coldef)
	termbox.SetCell(x+1, y, '┓', coldef, coldef)
	termbox.SetCell(x, y+1, '┗', coldef, coldef)
	termbox.SetCell(x+1, y+1, '┛', coldef, coldef)
	termbox.Flush()
}

func main() {
	print("Hello knnfiler")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	drawBox(0, 0)
MAINLOOP:
	for {
		w, h := termbox.Size()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			}
		}
		drawBox(rand.Intn(w), rand.Intn(h))
	}
}
