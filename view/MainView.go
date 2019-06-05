package view

import (
	"fmt"
	"regexp"
	"strings"
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
	drawX(0, 0, a.CurrentDirectory[a.Wid], data.COLOR_INDEX_WHITE, data.COLOR_INDEX_WHITE)
	drawX(0, 1, "         ", data.COLOR_INDEX_WHITE, data.COLOR_INDEX_BLACK)
}

func getRowColor(w int, i int) (int, int) {
	if w != a.Wid {
		return data.COLOR_INDEX_WHITE, data.COLOR_INDEX_WHITE
	}
	isActive := (i == a.CurrentCursorIndex[a.Wid])
	isDir := a.GetListFileInfo(w, i).IsDir
	if isActive {
		if isDir {
			return data.COLOR_INDEX_YELLOW, data.COLOR_INDEX_GREEN
		} else {
			return data.COLOR_INDEX_BLACK, data.COLOR_INDEX_GREEN
		}
	} else {
		if isDir {
			return data.COLOR_INDEX_YELLOW, data.COLOR_INDEX_WHITE
		} else {
			return data.COLOR_INDEX_WHITE, data.COLOR_INDEX_BLACK
		}
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
	drawX(1, 1, a.ConfirmMessage, data.COLOR_INDEX_WHITE, data.COLOR_INDEX_WHITE)
}
func search(){
	wid := a.Wid
	ll := a.FileListRowNum[wid]
	n := 1
	for i := 0; i < ll; i++ {
		fn := a.FileList[wid][i].FileName
		if strings.Contains(fn, a.SearchStr) || a.SearchStr == "" {
			if a.SearchCursorIndex == n {
				drawX(0, n, fn, data.COLOR_INDEX_BLACK, data.COLOR_INDEX_GREEN)
			} else {
				drawX(0, n, fn, 0, 0)
			}
			n += 1
		}
	}
	a.SearchHitNum = n
}
func textPreview(){
	drawX(0, 0, "Text preview", data.COLOR_INDEX_WHITE, data.COLOR_INDEX_BLACK)
	s := a.CurrentTargetContent
	for i, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(s, -1) {
		drawX(0, i+1, v, data.COLOR_INDEX_WHITE, data.COLOR_INDEX_BLACK)
    }
}
func Redraw(appData *data.AppData) {
	a = appData
	termbox.Clear(coldef, coldef)
	if a.WindowMode == data.WM_CONFIRM {
		drawConfirm()
	} else if a.WindowMode == data.WM_TEXT_PREVIEW {
		textPreview()
	} else if a.WindowMode == data.WM_SEARCH {
		search()
	} else {
		drawHeader()
		drawList(0)
		drawList(1)
	}
	termbox.Flush()
}
