package internal

import "fmt"

func Run() {
	InitGui()
	err := InitImg()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	GetMyGui().SetMine()
	GetMyGui().SetTool()
	GetMyGui().SetGui()
	GetMyGui().MainWindow.ShowAndRun()
}
