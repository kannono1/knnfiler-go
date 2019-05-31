package main

import (
	"./data"
	"./util"
	"./view"
	"github.com/nsf/termbox-go"
	"log"
)

var a = &data.AppData{}

func finalize() {
	a.Finalize()
}

func initialize() {
	// logfile, err := os.Create("test.log")
	// if err != nil {
	// 	panic("cannnot open test.log:" + err.Error())
	// }
	log.SetOutput(util.CreateFile(data.LOG_FILE))
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
		ev := termbox.PollEvent()
		if a.WindowMode == data.WM_FILER {
			switch ev.Type {
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
						a.Enter(a.Wid)
					case 113: // q
						break MAINLOOP
					case data.KEYCODE_x:
						a.Execute(a.Wid)
					}
				}
			}
		} else if a.WindowMode == data.WM_CONFIRM {
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					a.Escape()
				default:
					switch ev.Ch {
					case 121: // y
						if( a.WindowMode == data.WM_CONFIRM ) {
							a.Confirmed()
						}
					}
				}
			}
		} else {
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					a.Escape()
				default:
					switch ev.Ch {
					case 113: // q
						break MAINLOOP
					default:
						a.Escape()
					}
				}
			}
		}
		view.Redraw(a)
	}
	finalize()
}
