package main

import (
	"./data"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"log"
	"os"
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
var a = &data.AppData{}

func drawX(x, y int, str string, fgColor int, bgColor int) {
	runes := []rune(str)
	for _, r := range runes {
		termbox.SetCell(x, y, r, colors[fgColor], colors[bgColor])
		x += runewidth.RuneWidth(r)
	}
}

func drawHeader() {
	drawX(0, 0, a.CurrentDirectory[a.Wid], 0, 0)
	drawX(0, 1, "         ", 0, 1)
}

func getRowColor(w int, i int) (int, int) {
	if w != a.Wid {
		return 0, 0
	}
	if i == a.CurrentCursorIndex[a.Wid] {
		return 1, 3
	} else {
		return 0, 0
	}
}

func drawList(wid int) {
	w, h := termbox.Size()
	a.MaxScreenListRowNum = h - 2
	w2x := 0
	if wid == 1 {
		w2x = int(w / 2)
	}
	offset := a.CurrentCursorIndex[wid] - (a.MaxScreenListRowNum - 1) + (a.MaxScreenListRowNum - 1 - a.CurrentScreenCursorIndex[wid])
	ll := a.FileListRowNum[wid]
	if ll > a.MaxScreenListRowNum {
		ll = a.MaxScreenListRowNum
	}
	log.Print("Offset=", offset, a.CurrentCursorIndex[wid], a.MaxScreenListRowNum, " w=", wid, " si=", a.CurrentScreenCursorIndex[wid])
	for i := 0; i < ll; i++ {
		cf, cb := getRowColor(wid, i+offset)
		drawX(w2x, 2+i, a.FileList[wid][i+offset].FileName, cf, cb)
	}
}
func redraw() {
	termbox.Clear(coldef, coldef)
	drawHeader()
	drawList(0)
	drawList(1)
	termbox.Flush()
}

func initialize() {
	logfile, err := os.Create("test.log")
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	// defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("START !!")
	a.Initialize()
}

func switchWindow() {
	if a.Wid == 0 {
		a.Wid = 1
	} else {
		a.Wid = 0
	}
}
func enter() {
	a.EnterDir(a.Wid)
}

func main() {
	initialize()
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
			case termbox.KeyArrowDown:
				a.DownCursor(a.Wid)
			case termbox.KeyArrowUp:
				a.UpCursor(a.Wid)
			case termbox.KeyEsc:
				break MAINLOOP
			case termbox.KeyTab:
				switchWindow()
			default:
				//log.Print("key", ev.Key, ev.Ch)
				switch ev.Ch {
				case 99: // c
					a.Copy()
				case 104: // h
					a.GotoParentDir(a.Wid)
				case 106: // j
					a.DownCursor(a.Wid)
				case 107: // h
					a.UpCursor(a.Wid)
				case 108: // l
					enter()
				case 113: // q
					break MAINLOOP
				}
			}
		}
		redraw()
	}
}
