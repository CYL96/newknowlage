package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var RunGame *MyGuiExt

type MyGuiExt struct {
	app        fyne.App
	MainWindow fyne.Window
	Area       [][]BoomItemExt
	ToolBt     [4]*widget.Button // 0:reset, 1:open, 2:flog, 3:Exit
	Input      *widget.Entry
	TotalBoom  int
	LeftItem   int
	Size       int
}

func GetMyGui() *MyGuiExt {
	if RunGame == nil {
		RunGame = &MyGuiExt{}
	}
	return RunGame
}

func InitGui() {
	RunGame = new(MyGuiExt)
	RunGame.Size = 10
	RunGame.LeftItem = 10 * 10
	RunGame.app = app.New()
	RunGame.app.SetIcon(Img_mine)
	RunGame.MainWindow = RunGame.app.NewWindow("Boom")
	RunGame.MainWindow.SetFullScreen(false)
	RunGame.MainWindow.SetIcon(Img_mine)
	return
}

type BoomItemExt struct {
	Bt    *widget.Button
	Index int //-1代表雷  0-9 代表周围的雷数
	Show  int //  0代表没点击过 1代表已经点击过  2代表标志位雷 3代表？
}
