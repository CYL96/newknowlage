package internal

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	var (
		err  error
		icon fyne.Resource
	)
	err = os.Setenv("FYNE_FONT", "./config/wryh.ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Unsetenv("FYNE_FONT")
	ap := app.New()
	icon, err = fyne.LoadResourceFromPath("./config/icon.jpg")
	if err != nil {
		return
	}
	ap.SetIcon(icon)

	win := ap.NewWindow("Boom GUI")
	vbox := container.NewVBox()
	vbox2 := container.NewVBox()
	var resetbt *widget.Button
	resetbt = widget.NewButton("Reset", func() {
		vbox.Objects = []fyne.CanvasObject{}
		for i := 0; i < 10; i++ {
			hbox := container.NewHBox()
			for j := 0; j < 10; j++ {
				var bt *widget.Button
				bt = widget.NewButton(fmt.Sprintf("%d-%d", i, j), func() {
					fmt.Println(bt.Text)
					bt.SetText("233")
					bt.Disable()
				})
				hbox.Add(bt)
			}
			vbox.Add(hbox)
		}
	})
	vbox2.Add(resetbt)
	for i := 0; i < 10; i++ {
		hbox := container.NewHBox()
		for j := 0; j < 10; j++ {
			var bt *widget.Button
			bt = widget.NewButton(fmt.Sprintf("%d-%d", i, j), func() {
				bt.SetText("233")
				bt.Disable()
			})
			hbox.Add(bt)
		}
		vbox.Add(hbox)
	}
	vbox2.Add(vbox)
	win.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1), vbox2))
	win.ShowAndRun()
}
