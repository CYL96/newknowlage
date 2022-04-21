package internal

import "fyne.io/fyne/v2"

type MyGuiExt struct {
	app        fyne.App
	MainWindow fyne.Window
}

func InitGui() *MyGuiExt {
	return &MyGuiExt{}
}
