package data

type WindowMode int
const (
    WM_UNKNOWN WindowMode = iota
    WM_FILER
    WM_CONFIRM
    WM_TEXT_PREVIEW
    WM_SEARCH
)
