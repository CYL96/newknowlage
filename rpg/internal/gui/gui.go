package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MyGuiExt struct {
	app        fyne.App
	MainWindow fyne.Window
}

var Gui *MyGuiExt

func InitGui() (err error) {
	var (
		icon fyne.Resource
	)
	Gui = new(MyGuiExt)
	Gui.app = app.New()

	// 初始化图片
	icon, err = fyne.LoadResourceFromPath("./config/icon.jpg")
	if err != nil {
		return
	}
	Gui.app.SetIcon(icon)

	Gui.MainWindow = Gui.app.NewWindow("kill five")
	// Gui.MainWindow.SetFullScreen(true)
	Gui.MainWindow.Resize(fyne.NewSize(800, 600))
	Gui.MainWindow.SetFixedSize(true)
	RunUI(1)

	return
}

func GetMyGui() *MyGuiExt {
	return Gui
}
func (this *MyGuiExt) Run() {
	this.MainWindow.ShowAndRun()
}

func (this *MyGuiExt) SetStatus(txt string) {
}
