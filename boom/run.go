package boom

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"os"
)

func Run_Mine() {
	var ok bool
	tk := time.NewTicker(100 * time.Microsecond)
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	SetSizeAndNum(10)
	RunConfig.log, _ = os.OpenFile("boom.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 066)
	RunConfig.Key = make(chan termbox.Key)
	InitConfig()
	go KeyEvent()
	Printf_All()
	var key termbox.Key
	fmt.Println()
	ok = true
	RunConfig.StartTime = time.Now().Unix()
	for {
		select {
		case <-tk.C:
			Printf_All()
		case key = <-RunConfig.Key:
			switch key {
			case termbox.KeyEsc:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				termbox.Flush()
				termbox.Close()
				fmt.Println("exec game")
				return
			case termbox.KeyArrowUp:
				if RunConfig.y != 0 {
					RunConfig.y--
				}
				ok = true
			case termbox.KeyArrowDown:
				if RunConfig.y != RunConfig.Size-1 {
					RunConfig.y++
				}
				ok = true
			case termbox.KeyArrowRight:
				if RunConfig.x != RunConfig.Size-1 {
					RunConfig.x++
				}
				ok = true
			case termbox.KeyArrowLeft:
				if RunConfig.x != 0 {
					RunConfig.x--
				}
				ok = true
			case termbox.KeyF1:
				ok = SetFlag(RunConfig.x, RunConfig.y, reset)
			case termbox.KeyF2:
				ok = SetFlag(RunConfig.x, RunConfig.y, open)
			case termbox.KeyF3:
				ok = SetFlag(RunConfig.x, RunConfig.y, isMine)
			case termbox.KeyF4:
				ok = SetFlag(RunConfig.x, RunConfig.y, unknow)
			case termbox.KeySpace:
				InitConfig()
				ok = true
			case termbox.KeyPgup:
				RunConfig.Size += 5
				InitConfig()
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			case termbox.KeyPgdn:
				RunConfig.Size -= 5
				InitConfig()
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			default:
			}

		}
		if !ok {
			//fmt.Println(RunConfig.Start)
			GameOver()
			RunConfig.Start = 1
		}
		if RunConfig.Start == 0 {
			RunConfig.UseTime = time.Now().Unix() - RunConfig.StartTime
		}
		CheckWin()

	}
}

func KeyEvent() {
	for {
		RunConfig.Key <- termbox.PollEvent().Key
	}
}

func InitConfig() {
	RunConfig.Start = 0
	RunConfig.StartTime = time.Now().Unix()
	RunConfig.UseTime = 0
	RunConfig.Opened = 0
	RunConfig.TrueMine = 0
	RunConfig.IsMine = 0
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	SetMine()
	SetMineNum()
}

func CheckWin() {
	if RunConfig.Opened+RunConfig.MineNum == RunConfig.Size*RunConfig.Size ||
		RunConfig.TrueMine == RunConfig.MineNum {
		RunConfig.Start = 2
	}
}
