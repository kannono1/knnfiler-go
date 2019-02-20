package filesys
import (
	"io"
	"os"
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
    if  err != nil {
        panic(err)
    }
}