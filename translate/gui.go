package translate

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

func RunGui() {
	var inTE, outTE *walk.TextEdit
	declarative.MainWindow{
		Title:   "翻译",
		MinSize: declarative.Size{600, 400},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				AssignTo: &inTE,
				Name:     "源字符串",
				MinSize: struct {
					Width  int
					Height int
				}{Width: 100, Height: 100},
			},
			declarative.TextEdit{
				AssignTo: &outTE,
				Font: declarative.Font{
					PointSize: 14,
				},
				MinSize: struct {
					Width  int
					Height int
				}{Width: 200, Height: 300},
				ReadOnly: true,
			},
			declarative.PushButton{
				Text: "翻译",
				MinSize: struct {
					Width  int
					Height int
				}{Width: 200, Height: 100},
				OnClicked: func() {
					outTE.SetText(Translate(inTE.Text()))
				},
			},
		},
		ToolBar: declarative.ToolBar{},
	}.Run()
}
