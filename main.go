package main

import (
	"./data"
	"./view"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

var a = &data.AppData{}

func initialize() {
	logfile, err := os.Create("test.log")
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	log.SetOutput(logfile)
	log.Println("START !!")
	a.Initialize()
}

func main() {
	initialize()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	view.Redraw(a)
MAINLOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				a.DownCursor(a.Wid)
			case termbox.KeyArrowUp:
				a.UpCursor(a.Wid)
			case termbox.KeyEsc:
				a.Escape()
			case termbox.KeyTab:
				a.SwitchWindow()
			default:
				//log.Print("key", ev.Key, ev.Ch)
				switch ev.Ch {
				case 99: // c
					a.Copy()
				case 100: // d
					a.DeleteConfirm()
				case 104: // h
					a.GotoParentDir(a.Wid)
				case 106: // j
					a.DownCursor(a.Wid)
				case 107: // h
					a.UpCursor(a.Wid)
				case 108: // l
					a.EnterDir(a.Wid)
				case 113: // q
					break MAINLOOP
				}
			}
		}
		view.Redraw(a)
	}
}
