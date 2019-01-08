package data

import (
	"strconv"
)

type AppData struct{
	CurrentDirectory[2]string
	CurrentCursorIndex[2]int
	FileList[2][]FileInfo
}

func (a *AppData) Initialize() {
	a.CurrentDirectory[0] = "aaaa"
	a.CurrentDirectory[1] = "bbbb"
	a.FileList[0] = make([]FileInfo, 3)
    for i, _ := range a.FileList[0] {
		a.FileList[0][i].FileName = "FILE" + strconv.Itoa(i)
	}
}