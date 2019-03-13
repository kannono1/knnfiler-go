package view

import (
	"../data"
	"strconv"
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
var a *data.AppData

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
	w2x := int(w / 2)
	left:= 0
	if wid == 1 {
		left = w2x 
	}
	offset := a.CurrentCursorIndex[wid] - (a.MaxScreenListRowNum - 1) + (a.MaxScreenListRowNum - 1 - a.CurrentScreenCursorIndex[wid])
	ll := a.FileListRowNum[wid]
	if ll > a.MaxScreenListRowNum {
		ll = a.MaxScreenListRowNum
	}
	for i := 0; i < ll; i++ {
		cf, cb := getRowColor(wid, i+offset)
		drawX(left, 2+i, a.FileList[wid][i+offset].FileName, cf, cb)
		drawX(left+w2x-8, 2+i, strconv.FormatInt(a.FileList[wid][i+offset].FileSize, 10), cf, cb)
	}
}
func drawConfirm(){
	drawX(1, 1, a.ConfirmMessage, 0, 0)
}
func Redraw(appData *data.AppData) {
	a = appData
	termbox.Clear(coldef, coldef)
	if a.WindowMode == data.WM_CONFIRM {
		drawConfirm()
	} else {
		drawHeader()
		drawList(0)
		drawList(1)
	}
	termbox.Flush()
}
