package main

import (
	"./data"
    // "log"
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

func drawList() {
	w, _ := termbox.Size()
	w2x := int(w / 2)
	for i, _ := range a.FileList[0] {
		cf, cb := getRowColor(0, i)
		drawX(0, 2+i, a.FileList[0][i].FileName, cf, cb)
	}
	for i, _ := range a.FileList[1] {
		cf, cb := getRowColor(1, i)
		drawX(w2x, 2+i, a.FileList[1][i].FileName, cf, cb)
	}
}

func redraw() {
	termbox.Clear(coldef, coldef)
	drawHeader()
	drawList()
	termbox.Flush()
}

func initialize() {
	a.Initialize()
}

func cursorDown() {
	if a.CurrentCursorIndex[a.Wid] < a.FileListRowNum[a.Wid]-1 {
		a.CurrentCursorIndex[a.Wid] += 1
	}
}
func cursorUp() {
	if a.CurrentCursorIndex[a.Wid] > 0 {
		a.CurrentCursorIndex[a.Wid] -= 1
	}
}
func switchWindow(){
    if(a.Wid == 0){
        a.Wid = 1
    } else {
        a.Wid = 0
    }
}
func enter(){
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
				cursorDown()
			case termbox.KeyArrowUp:
				cursorUp()
			case termbox.KeyEsc:
				break MAINLOOP
            case termbox.KeyTab:
                switchWindow()
			default:
				//log.Print("key", ev.Key, ev.Ch)
				switch ev.Ch {
                case 104: // h
                    a.GotoParentDir(a.Wid)
				case 106: // j
					cursorDown()
				case 107: // h
					cursorUp()
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
