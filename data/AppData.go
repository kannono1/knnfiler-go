package data

import (
	"strconv"
)

type AppData struct{
	Wid int
	CurrentDirectory[2]string
	CurrentCursorIndex[2]int
	FileList[2][]FileInfo
}

func (a *AppData) Initialize() {
	a.CurrentDirectory[0] = "aaaa"
	a.CurrentDirectory[1] = "bbbb"
	a.FileList[a.Wid] = make([]FileInfo, 13)
    for i, _ := range a.FileList[a.Wid] {
		a.FileList[a.Wid][i].FileName = "FILE" + strconv.Itoa(i)
	}
}