package gui

import (
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func MainUI() {
	User.Rand()
	Monster.Rand()
	gui := GetMyGui()
	// gui.MainWindow.Resize(fyne.Size{
	// 	Width:  600,
	// 	Height: 500,
	// })
	user := User.BindUi()
	monster := Monster.BindUi()

	// status := widget.NewMultiLineEntry()
	// status.Wrapping = fyne.TextWrapWord
	// status.Disable()
	btt := container.NewVBox()
	round := widget.NewLabelWithData(binding.IntToStringWithFormat(User.LevelBd, "ROUND: %d"))
	btt.Add(round)
	start := widget.NewButton("Start", func() {
		AttackUI()
	})
	btt.Add(start)
	Flush := widget.NewButton("Rebuild", func() {
		User.Rand()
		Monster.Rand()
		// MainUI()
	})
	btt.Add(Flush)
	exit := widget.NewButton("Exit", func() {
		os.Exit(1)
	})
	btt.Add(exit)

	right := container.NewHSplit(
		user,
		btt,
	)
	left := container.NewHSplit(
		right,
		monster,
	)

	gui.MainWindow.SetContent(left)
}
func AttackUI() {

	gui := GetMyGui()
	// gui.MainWindow.Resize(fyne.Size{
	// 	Width:  600,
	// 	Height: 500,
	// })
	user := User.BindUi()
	monster := Monster.BindUi()

	btt := container.NewPadded()
	round := widget.NewLabelWithData(binding.IntToStringWithFormat(User.LevelBd, "ROUND: %d"))
	btt.Add(round)

	list := User.SetAttackUi()
	vb := container.NewVBox()
	for i := range list {
		vb.Add(list[i])
	}
	status := container.NewVScroll(vb)
	btt.Add(status)

	right := container.NewHSplit(
		user,
		btt,
	)
	left := container.NewHSplit(
		right,
		monster,
	)

	gui.MainWindow.SetContent(left)
	ok := Attack()
	if ok {
		UpgradeUI()
	} else {
		dialog.ShowInformation("you five ", "you get round:"+strconv.Itoa(User.Level), gui.MainWindow)
		MainUI()
	}
}
func UpgradeUI() {

	gui := GetMyGui()
	// gui.MainWindow.Resize(fyne.Size{
	// 	Width:  600,
	// 	Height: 500,
	// })
	user := User.BindUi()
	monster := Monster.BindUi()

	btt := container.NewVBox()
	round := widget.NewLabelWithData(binding.IntToStringWithFormat(User.LevelBd, "ROUND: %d"))
	btt.Add(round)

	detail := User.Upgrade()
	btt.Add(detail)

	right := container.NewHSplit(
		user,
		btt,
	)
	left := container.NewHSplit(
		right,
		monster,
	)
	gui.MainWindow.SetContent(left)
}

func TargetUi() fyne.CanvasObject {
	vb := container.NewVBox()
	hb := container.NewHBox()

	name := widget.NewLabel("userName")
	vb.Add(name)
	dd := 0.10

	bar := widget.NewProgressBarWithData(binding.BindFloat(&dd))
	vb.Add(bar)

	hp := widget.NewLabel("HP:100")
	vb.Add(hp)
	att := widget.NewLabel("ATT:100")
	vb.Add(att)

	a := widget.NewLabel("A")
	hb.Add(a)
	b := widget.NewLabel("B")
	hb.Add(b)
	c := widget.NewLabel("C")
	hb.Add(c)

	vb.Add(hb)
	return vb
}
