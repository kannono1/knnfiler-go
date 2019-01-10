package data

import (
	"strconv"
)

type AppData struct{
	Wid int
	CurrentDirectory[2]string
	CurrentCursorIndex[2]int
	FileList[2][]FileInfo
	FileListRowNum[2]int
}

func (a *AppData) Initialize() {
	a.CurrentDirectory[0] = "aaaa"
	a.CurrentDirectory[1] = "bbbb"
	a.FileList[0] = make([]FileInfo, 13)
	a.FileList[1] = make([]FileInfo, 11)
    for i, _ := range a.FileList[0] {
		a.FileList[0][i].FileName = "FILEA" + strconv.Itoa(i)
	}
    for i, _ := range a.FileList[1] {
		a.FileList[1][i].FileName = "FILEB" + strconv.Itoa(i)
	}
	a.FileListRowNum[0] = len(a.FileList[0])
	a.FileListRowNum[1] = len(a.FileList[1])
}