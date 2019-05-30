package data
import (
	"strconv"
	"time"
)
type FileInfo struct{
	FileName string
	FileSize int64
	IsDir bool
	ModTime time.Time
}
func (a FileInfo) Initialize() {
	a.FileName = "ffffffffffff"
}
func (a FileInfo) GetFileSizeStr() string {
	if a.IsDir {
		return "<DIR>"
	} else {
		return strconv.FormatInt(a.FileSize, 10)
	}
}
func (a FileInfo) GetModTimeStr() string {
	return a.ModTime.Format("2006-01-02 15:04")
}