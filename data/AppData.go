package data
import (
	"../filesys"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)
type AppData struct {
	Wid                      int
	ConfirmMessage string
	ConfirmedFunction func()
	CurrentDirectory         [2]string
	CurrentCursorIndex       [2]int
	CurrentScreenCursorIndex [2]int
	FileList                 [2][]FileInfo
	FileListRowNum           [2]int
	MaxScreenListRowNum      int
	WindowMode WindowMode
}
func (a *AppData) ReadDir(wid int, dir string) {
	log.Print("-- ReadDir ", dir);
	files, _ := ioutil.ReadDir(dir)
	a.FileListRowNum[wid] = len(files)
	a.FileList[wid] = make([]FileInfo, a.FileListRowNum[wid])
	for i, f := range files {
		a.FileList[wid][i].FileName = f.Name()
	}
	if a.CurrentCursorIndex[wid] >= len(files) {
		a.initCursorIndex(wid)
	}
}
func (a *AppData) EnterDir(wid int) {
	dir := filepath.Join(a.CurrentDirectory[wid], a.GetListFileName(wid, a.CurrentCursorIndex[wid]))
	a.GotoDir(wid, dir)
}
func (a *AppData) Copy() {
	cwid := a.Wid
	owid := a.Wid^1
	fn := a.GetListFileName(cwid, a.CurrentCursorIndex[cwid])
	from := filepath.Join(a.CurrentDirectory[cwid], fn)
	to   := filepath.Join(a.CurrentDirectory[owid], fn)
	filesys.Copy(from, to)
	a.ReadDir(owid, a.CurrentDirectory[owid])
}
func (a *AppData) Escape() {
	a.WindowMode = WM_FILER
}
func (a *AppData) DeleteConfirm() {
	a.WindowMode = WM_CONFIRM
	a.ConfirmMessage = "Are you sure you want to delete ?"
	a.ConfirmedFunction = a.Delete
}
func (a *AppData) Confirmed() {
	a.WindowMode = WM_FILER
	a.ConfirmedFunction()
}
func (a *AppData) Delete() {
	cwid := a.Wid
	fn := a.GetListFileName(cwid, a.CurrentCursorIndex[cwid])
	src := filepath.Join(a.CurrentDirectory[cwid], fn)
	filesys.Delete(src)
	a.ReadDir(cwid, a.CurrentDirectory[cwid])
}
func (a *AppData) GetListFileName(wid int, i int) string {
	return a.FileList[wid][i].FileName
}
func (a *AppData) initCursorIndex(wid int) {
	a.CurrentCursorIndex[wid] = 0
	a.CurrentScreenCursorIndex[wid] = 0
}
func (a *AppData) GotoDir(wid int, dir string) {
	a.initCursorIndex(wid)
	a.CurrentDirectory[wid] = dir
	a.ReadDir(wid, a.CurrentDirectory[wid])
}
func (a *AppData) GotoParentDir(wid int) {
	a.GotoDir(wid, filepath.Dir(a.CurrentDirectory[wid]))
}
func (a *AppData) DownCursor(wid int) {
	a.CurrentScreenCursorIndex[wid]++
	if a.CurrentScreenCursorIndex[wid] > (a.MaxScreenListRowNum - 1) {
		a.CurrentScreenCursorIndex[wid] = (a.MaxScreenListRowNum - 1)
	} else if a.CurrentScreenCursorIndex[wid] > (a.FileListRowNum[wid] - 1) {
		a.CurrentScreenCursorIndex[wid] = (a.FileListRowNum[wid] - 1)
	}
	if a.CurrentCursorIndex[a.Wid] < a.FileListRowNum[a.Wid]-1 {
		a.CurrentCursorIndex[a.Wid] += 1
	}
}
func (a *AppData) UpCursor(wid int) {
	a.CurrentScreenCursorIndex[wid]--
	if a.CurrentScreenCursorIndex[wid] < 0 {
		a.CurrentScreenCursorIndex[wid] = 0
	}
	if a.CurrentCursorIndex[a.Wid] > 0 {
		a.CurrentCursorIndex[a.Wid] -= 1
	}
}
func (a *AppData) Initialize() {
	a.CurrentDirectory[0], _ = os.Getwd()
	a.CurrentDirectory[1], _ = os.Getwd()
	a.ReadDir(0, a.CurrentDirectory[0])
	a.ReadDir(1, a.CurrentDirectory[1])
}
func (a *AppData) SwitchWindow() {
	if a.Wid == 0 {
		a.Wid = 1
	} else {
		a.Wid = 0
	}
}