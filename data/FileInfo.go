package data

type FileInfo struct{
	FileName string
	FileSize int64
}
func (a FileInfo) Initialize() {
	a.FileName = "ffffffffffff"
}