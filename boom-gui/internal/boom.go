package internal

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var Action int // 1:open  2:flag  3:question
func BoomButton(x, y int) func() {
	if x < 0 || y < 0 || x >= RunGame.Size || y >= RunGame.Size {
		return func() {}
	}
	return func() {
		switch Action {
		case 1:
			if RunGame.Area[x][y].Show == 2 {
				return
			}
			RunGame.Area[x][y].Show = 1
			if RunGame.Area[x][y].Index == -1 {
				RunGame.Area[x][y].Bt.SetIcon(Img_mine)
				RunGame.GameOver()
			} else {
				RunGame.OpenZeroItem(x, y)
			}
			RunGame.Area[x][y].Bt.Disable()
		case 2:
			if RunGame.Area[x][y].Show == 2 {
				RunGame.Area[x][y].Show = 0
				RunGame.Area[x][y].Bt.SetIcon(Img_invisible)
			} else {
				RunGame.Area[x][y].Show = 2
				RunGame.Area[x][y].Bt.SetIcon(Img_flag)
			}
		case 3:
			RunGame.Area[x][y].Show = 3
			RunGame.Area[x][y].Bt.SetIcon(Img_question)
		default:
			return
		}

		if RunGame.LeftItem == RunGame.TotalBoom {
			RunGame.GameSuccess()
		}
	}
}
func (this *MyGuiExt) GameSuccess() {

	Action = 0
	this.ToolBt[1].Disable()
	this.ToolBt[2].Disable()
	this.ToolBt[3].Disable()
	label := widget.NewLabel("you win")
	label.Alignment = fyne.TextAlignCenter
	dialog.NewCustom("information", "OK", label, this.MainWindow).Show()
}

func (this *MyGuiExt) GameOver() {
	Action = 0
	this.ToolBt[1].Disable()
	this.ToolBt[2].Disable()
	this.ToolBt[3].Disable()
	for i := range this.Area {
		for j := range this.Area[i] {
			this.Area[i][j].Bt.Disable()
			if this.Area[i][j].Index == -1 {
				this.Area[i][j].Bt.SetIcon(Img_mine)
			} else {
				this.Area[i][j].Bt.SetIcon(ImgNum[this.Area[i][j].Index])
			}
		}
	}
	label := widget.NewLabel("you lose")
	label.Alignment = fyne.TextAlignCenter
	dialog.NewCustom("information", "OK", label, this.MainWindow).Show()
}
func (this *MyGuiExt) Reset() {
	Action = 1
	this.ToolBt[1].Disable()
	this.ToolBt[2].Enable()
	this.ToolBt[3].Enable()
	if this.Size < 5 {
		this.Size = 5
	}
	if this.Size >= 30 {
		this.Size = 30
	}
	this.LeftItem = this.Size * this.Size
	this.SetMine()
	this.SetGui()
}
func (this *MyGuiExt) OpenZeroItem(x, y int) {
	this.Area[x][y].Show = 1
	this.Area[x][y].Bt.SetIcon(ImgNum[this.Area[x][y].Index])
	this.Area[x][y].Bt.Disable()
	this.LeftItem--
	if this.Area[x][y].Index == 0 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if (i == 0 && j == 0) || x+i < 0 || y+j < 0 || x+i >= this.Size || y+j >= this.Size {
					continue
				}
				if this.Area[x+i][y+j].Index == -1 {
					continue
				}
				if this.Area[x+i][y+j].Show == 0 {
					this.OpenZeroItem(x+i, y+j)
				}
			}
		}
	}
}
func (this *MyGuiExt) SetTool() {
	// reset
	this.ToolBt[0] = new(widget.Button)
	this.ToolBt[0].SetText("Reset")
	this.ToolBt[0].OnTapped = func() {
		this.Reset()
	}
	Action = 1
	this.ToolBt[1] = new(widget.Button)
	this.ToolBt[1].SetText("Open")
	this.ToolBt[1].Disable()
	this.ToolBt[1].OnTapped = func() {
		if Action == 2 {
			this.ToolBt[2].Enable()
		} else {
			this.ToolBt[3].Enable()
		}
		this.ToolBt[1].Disable()
		Action = 1
	}
	this.ToolBt[2] = new(widget.Button)
	this.ToolBt[2].SetText("Flag")
	this.ToolBt[2].OnTapped = func() {
		if Action == 1 {
			this.ToolBt[1].Enable()
		} else {
			this.ToolBt[3].Enable()
		}
		this.ToolBt[2].Disable()
		Action = 2
	}
	this.ToolBt[3] = new(widget.Button)
	this.ToolBt[3].SetText("Question")
	this.ToolBt[3].OnTapped = func() {
		if Action == 1 {
			this.ToolBt[1].Enable()
		} else {
			this.ToolBt[2].Enable()
		}
		this.ToolBt[3].Disable()
		Action = 3
	}
}

func (this *MyGuiExt) SetGui() {
	vbox := container.NewVBox()

	data := binding.BindInt(&this.Size)
	this.Input = widget.NewEntry()
	this.Input.Bind(binding.IntToString(data))

	toolBox := container.NewHBox(this.Input, RunGame.ToolBt[0], RunGame.ToolBt[1], RunGame.ToolBt[2], RunGame.ToolBt[3])
	vbox.Add(toolBox)

	for i := range this.Area {
		hbox := container.NewHBox()
		for j := range this.Area[i] {
			hbox.Add(this.Area[i][j].Bt)
		}
		vbox.Add(hbox)
	}
	this.MainWindow.SetContent(vbox)
}

func (this *MyGuiExt) SetMine() {
	if this.Size < 5 {
		this.Size = 5
	}
	this.TotalBoom = (this.Size-5)*3 + 5
	this.Area = make([][]BoomItemExt, this.Size)
	for i := range this.Area {
		this.Area[i] = make([]BoomItemExt, this.Size)
	}
	for i := range this.Area {
		for j := range this.Area[i] {
			// 	设置按钮
			this.Area[i][j].Bt = new(widget.Button)
			this.Area[i][j].Bt.SetIcon(Img_invisible)
			this.Area[i][j].Bt.SetText("")
			this.Area[i][j].Bt.OnTapped = BoomButton(i, j)
		}
	}
	data := make(map[string]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < this.TotalBoom; i++ {
	MINE:
		x := rand.Intn(this.Size)
		y := rand.Intn(this.Size)
		if _, ok := data[fmt.Sprintf("%d,%d", x, y)]; ok {
			goto MINE
		} else {
			this.Area[y][x].Index = -1
			data[fmt.Sprintf("%d,%d", x, y)] = 1
			// 	周围格子的值+1
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if x+j >= 0 && x+j < this.Size && y+k >= 0 && y+k < this.Size && this.Area[y+k][x+j].Index != -1 {
						this.Area[y+k][x+j].Index++
					}
				}
			}
		}
	}
}
