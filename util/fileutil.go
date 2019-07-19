package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"github.com/otiai10/copy"
)

func Copy(from string, to string) {
	log.Println("Copy: " + from + " to " + to)
	err := copy.Copy(from, to)
	if err != nil {
		panic(err)
	}
}
func CreateDir(dir string) {
    if err := os.Mkdir(dir, 0766); err != nil {
        log.Println("Exist Dir: " + dir)
    }
}
func CreateFile(path string) (*os.File) {
    f, err := os.Create(path) 
    if err != nil {
        log.Println(err)
    }
    return f
}
func Delete(src string) {
	if err := os.RemoveAll(src); err != nil {
		panic(err)
	}
	log.Println("-- Deleted ", src)
}
func Execute(path string) {
    if err := exec.Command("open", path).Start(); err != nil { // Mac
        panic(err)
    }
}
func ReadFile(path string) (string) {
	f, err := os.Open(path)
	if err != nil {
        log.Println("File read error")
        return ""
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b)
}
func WriteFile(path string, s string) {
    f := CreateFile(path)
    f.WriteString(s)
    f.Sync()
}
