package view

import (
	"fmt"
	"../data"
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

func getRowString(fn string, size string, modtime string) (string) {
	fn2 := runewidth.FillRight(fn, 30)
	fn3 := runewidth.Truncate(fn2, 30, "...")
	size2 := runewidth.FillLeft(size, 12)
	return fmt.Sprintf("%s %s %s", fn3, size2, modtime)
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
		fn := a.FileList[wid][i+offset].FileName
		size := a.FileList[wid][i+offset].GetFileSizeStr()
		modtime := a.FileList[wid][i+offset].GetModTimeStr()
		s := getRowString(fn, size, modtime)
		drawX(left, 2+i, s, cf, cb)
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
