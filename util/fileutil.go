package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Copy(from string, to string) {
	src, err := os.Open(from)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.Create(to)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
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
