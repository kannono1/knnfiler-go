package data
import (
	"strconv"
)
type FileInfo struct{
	FileName string
	FileSize int64
	IsDir bool
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