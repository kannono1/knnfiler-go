package data

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type AppData struct{
	Wid int
	CurrentDirectory[2]string
	CurrentCursorIndex[2]int
	FileList[2][]FileInfo
	FileListRowNum[2]int
}

func (a *AppData) ReadDir(wid int, dir string) {
	files, _ := ioutil.ReadDir(dir)
	a.FileListRowNum[wid] = len(files)
	a.FileList[wid] = make([]FileInfo, a.FileListRowNum[wid])
    for i, f := range files {
		a.FileList[wid][i].FileName = f.Name()
	}
}
func (a *AppData) GotoParentDir(wid int){
	a.CurrentCursorIndex[wid] = 0
	a.CurrentDirectory[wid] = filepath.Dir(a.CurrentDirectory[wid])
	a.ReadDir(wid, a.CurrentDirectory[wid])
}

func (a *AppData) Initialize() {
	a.CurrentDirectory[0], _ = os.Getwd()
	a.CurrentDirectory[1], _ = os.Getwd()
	a.ReadDir(0, a.CurrentDirectory[0])
	a.ReadDir(1, a.CurrentDirectory[1])
}