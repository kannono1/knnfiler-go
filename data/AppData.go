package data
import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"../util"
	"github.com/mattn/go-runewidth"
)
type AppData struct {
	Wid                      int
	ConfirmMessage string
	ConfirmedFunction func()
	CurrentDirectory         [2]string
	CurrentCursorIndex       [2]int
	CurrentScreenCursorIndex [2]int
	CurrentTargetContent string
	FileList                 [2][]FileInfo
	FileListRowNum           [2]int
	MaxScreenListRowNum      int
	SearchCursorIndex int
	SearchHitNum int
	SearchStr string
	WindowMode WindowMode
}
func (a *AppData) Confirmed() {
	a.WindowMode = WM_FILER
	a.ConfirmedFunction()
}
func (a *AppData) Copy() {
	cwid := a.Wid
	owid := a.Wid^1
	fn := a.GetListFileName(cwid, a.CurrentCursorIndex[cwid])
	from := filepath.Join(a.CurrentDirectory[cwid], fn)
	to   := filepath.Join(a.CurrentDirectory[owid], fn)
	util.Copy(from, to)
	a.ReadDir(owid, a.CurrentDirectory[owid])
}
func (a *AppData) DeleteConfirm() {
	a.WindowMode = WM_CONFIRM
	a.ConfirmMessage = "Are you sure you want to delete ?"
	a.ConfirmedFunction = a.Delete
}
func (a *AppData) Delete() {
	cwid := a.Wid
	fn := a.GetListFileName(cwid, a.CurrentCursorIndex[cwid])
	src := filepath.Join(a.CurrentDirectory[cwid], fn)
	util.Delete(src)
	a.ReadDir(cwid, a.CurrentDirectory[cwid])
}
func (a *AppData) DownCursor(wid int, v int) {
	a.CurrentScreenCursorIndex[wid] += v
	if a.CurrentScreenCursorIndex[wid] > (a.MaxScreenListRowNum - 1) {
		a.CurrentScreenCursorIndex[wid] = (a.MaxScreenListRowNum - 1)
	} else if a.CurrentScreenCursorIndex[wid] > (a.FileListRowNum[wid] - 1) {
		a.CurrentScreenCursorIndex[wid] = (a.FileListRowNum[wid] - 1)
	}
	if a.CurrentCursorIndex[a.Wid] < a.FileListRowNum[a.Wid] -v {
		a.CurrentCursorIndex[a.Wid] += v
	}
}
func (a *AppData) Enter(wid int) {
	ind := a.CurrentCursorIndex[wid]
	isDir := a.GetListFileInfo(wid, ind).IsDir
	path := filepath.Join(a.CurrentDirectory[wid], a.GetListFileName(wid, ind))
	if isDir {
		a.GotoDir(wid, path)
	} else {
		a.Preview(wid, path)
	}
}
func (a *AppData) Execute(wid int) {
	ind := a.CurrentCursorIndex[wid]
	// isDir := a.GetListFileInfo(wid, ind).IsDir
	path := filepath.Join(a.CurrentDirectory[wid], a.GetListFileName(wid, ind))
	util.Execute(path)
}
func (a *AppData) Escape() {
	a.WindowMode = WM_FILER
}
func (a *AppData) Finalize() {
	wid := 0
	dir := a.CurrentDirectory[wid]
	util.WriteFile(HISTORY_FILE_0, dir)
}
func (a *AppData) GetListFileInfo(wid int, i int) FileInfo {
	return a.FileList[wid][i]
}
func (a *AppData) GetListFileName(wid int, i int) string {
	return a.FileList[wid][i].FileName
}
func (a *AppData) GotoDir(wid int, dir string) {
	a.initCursorIndex(wid)
	a.CurrentDirectory[wid] = dir
	a.ReadDir(wid, a.CurrentDirectory[wid])
}
func (a *AppData) GotoParentDir(wid int) {
	a.GotoDir(wid, filepath.Dir(a.CurrentDirectory[wid]))
}
func (a *AppData) initCursorIndex(wid int) {
	a.CurrentCursorIndex[wid] = 0
	a.CurrentScreenCursorIndex[wid] = 0
}
func (a *AppData) Initialize() {
	a.WindowMode = WM_FILER
	util.CreateDir(STORAGE_DIR)
	dir := util.ReadFile(HISTORY_FILE_0)
	if dir == "" {
		dir, _ = os.Getwd()
	}
	a.CurrentDirectory[0] = dir
	a.CurrentDirectory[1], _ = os.Getwd()
	a.ReadDir(0, a.CurrentDirectory[0])
	a.ReadDir(1, a.CurrentDirectory[1])
}
func (a *AppData) Preview(wid int, path string) {
	a.WindowMode = WM_TEXT_PREVIEW
	a.ReadFileForPreview(wid, path)
}
func (a *AppData) ReadFileForPreview(wid int, path string) {
	a.CurrentTargetContent = util.TabToSpace( util.ReadFile(path) )
}
func (a *AppData) ReadDir(wid int, dir string) {
	log.Print("-- ReadDir ", dir)
	files, _ := ioutil.ReadDir(dir)
	a.FileListRowNum[wid] = len(files)
	a.FileList[wid] = make([]FileInfo, a.FileListRowNum[wid])
	for i, f := range files {
		a.FileList[wid][i].FileName = f.Name()
		a.FileList[wid][i].FileSize = f.Size()
		a.FileList[wid][i].IsDir = f.IsDir()
		a.FileList[wid][i].ModTime = f.ModTime()
	}
	if a.CurrentCursorIndex[wid] >= len(files) {
		a.initCursorIndex(wid)
	}
}
func (a *AppData) SearchAddString(s string) {
	a.SearchStr += s
}
func (a *AppData) SearchDeleteString() {
	c := runewidth.StringWidth(a.SearchStr)
	if c > 1 {
		a.SearchStr = a.SearchStr[0:c-1]
	} else {
		a.SearchStr = ""
	}
}
func (a *AppData) SearchCursorDown() {
	a.SearchCursorIndex += 1
	if a.SearchCursorIndex > a.SearchHitNum -1 {
		a.SearchCursorIndex = a.SearchHitNum -1
	}
}
func (a *AppData) SearchCursorUp() {
	a.SearchCursorIndex -= 1
	if a.SearchCursorIndex < 1 {
		a.SearchCursorIndex = 1
	}
}
func (a *AppData) SearchEnter() {
	a.WindowMode = WM_FILER
	wid := a.Wid
	ll := a.FileListRowNum[wid]
	n := 1
	v := 0
	for i := 0; i < ll; i++ {
		fn := a.FileList[wid][i].FileName
		if strings.Contains(fn, a.SearchStr) || a.SearchStr == "" {
			if a.SearchCursorIndex == n {
				v = i
				break
			}
			n += 1
		}
	}
	log.Print("SearchEnter:", v )
	a.CurrentCursorIndex[wid] = 0
	a.CurrentScreenCursorIndex[wid] = 0
	a.DownCursor(wid, v)
}
func (a *AppData) SearchStart(wid int) {
	a.WindowMode = WM_SEARCH
	a.SearchCursorIndex = 1
	a.SearchStr = ""
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
func (a *AppData) SwitchWindow() {
	if a.Wid == 0 {
		a.Wid = 1
	} else {
		a.Wid = 0
	}
}