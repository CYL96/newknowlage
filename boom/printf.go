package boom

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"time"
)

func PrintfMine1() {
	for y := range RunConfig.Area {
		for x := range RunConfig.Area[y] {
			switch RunConfig.Area[y][x].Open {
			case reset:
				fmt.Print("[ ]")
			case open:
				if RunConfig.Area[y][x].Mine {
					fmt.Print("[#]")
				} else {
					fmt.Printf("[%d]", RunConfig.Area[y][x].MineNum)
				}
			case isMine:
				fmt.Print("[M]")
			case unknow:
				fmt.Print("[?]")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func Printf_All() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	PrintfMine()
	Printf_Message()
	//PrintfMine()
	termbox.Flush()

}

func PrintfMine() {
	bg := termbox.ColorBlack
	for y := range RunConfig.Area {
		for x := range RunConfig.Area[y] {
			if RunConfig.x == x && RunConfig.y == y {
				bg = termbox.ColorYellow
			} else {
				bg = termbox.ColorBlack
			}
			switch RunConfig.Area[y][x].Open {
			case reset:
				termbox.SetCell(3*x, y, '[', termbox.ColorWhite, bg)
				termbox.SetCell(3*x+1, y, '/', termbox.ColorWhite, bg)
				termbox.SetCell(3*x+2, y, ']', termbox.ColorWhite, bg)
			case open:
				if RunConfig.Area[y][x].Mine {
					termbox.SetCell(3*x, y, '[', termbox.ColorRed, bg)
					termbox.SetCell(3*x+1, y, '#', termbox.ColorRed, bg)
					termbox.SetCell(3*x+2, y, ']', termbox.ColorRed, bg)
				} else if RunConfig.Area[y][x].MineNum-48 == 0 {
					termbox.SetCell(3*x, y, '[', termbox.ColorGreen, bg)
					termbox.SetCell(3*x+1, y, ' ', termbox.ColorGreen, bg)
					termbox.SetCell(3*x+2, y, ']', termbox.ColorGreen, bg)
				} else {
					termbox.SetCell(3*x, y, '[', termbox.ColorGreen, bg)
					termbox.SetCell(3*x+1, y, int32(RunConfig.Area[y][x].MineNum), termbox.ColorGreen, bg)
					termbox.SetCell(3*x+2, y, ']', termbox.ColorGreen, bg)
				}
			case isMine:
				termbox.SetCell(3*x, y, '[', termbox.ColorMagenta, bg)
				termbox.SetCell(3*x+1, y, 'M', termbox.ColorMagenta, bg)
				termbox.SetCell(3*x+2, y, ']', termbox.ColorMagenta, bg)
			case unknow:
				termbox.SetCell(3*x, y, '[', termbox.ColorBlue, bg)
				termbox.SetCell(3*x+1, y, '?', termbox.ColorBlue, bg)
				termbox.SetCell(3*x+2, y, ']', termbox.ColorBlue, bg)
			}
		}
	}
}

func Printf_Message() {
	printf_String(3*RunConfig.Size+10, 0, time.Now().Format("Time: 2006-01-02 15:04:05"))
	printf_String(3*RunConfig.Size+10, 1, fmt.Sprintf("UseTime: %d:%02d:%02d", RunConfig.UseTime/3600, RunConfig.UseTime/60, RunConfig.UseTime%60))
	printf_String(3*RunConfig.Size+10, 2, fmt.Sprintf("Size: %d * %d", RunConfig.Size, RunConfig.Size))
	printf_String(3*RunConfig.Size+10, 3, fmt.Sprintf("Open: %d ; Unopend: %d", RunConfig.Opened, RunConfig.Size*RunConfig.Size-RunConfig.Opened))
	printf_String(3*RunConfig.Size+10, 4, fmt.Sprintf("Boom Num: %d", RunConfig.MineNum))
	printf_String(3*RunConfig.Size+10, 5, fmt.Sprintf("Sign Boom Num: %d", RunConfig.IsMine))
	if RunConfig.Start == 1 {
		printf_String_Color(3*RunConfig.Size+10, 7, termbox.ColorRed, termbox.ColorDefault, "GAME OVER")
		printf_String_Color(3*RunConfig.Size+10, 8, termbox.ColorRed, termbox.ColorDefault, "SPACE TO RESTART")
	} else if RunConfig.Start == 2 {
		printf_String_Color(3*RunConfig.Size+10, 7, termbox.ColorRed, termbox.ColorDefault, "YOU WIN")
		printf_String_Color(3*RunConfig.Size+10, 8, termbox.ColorRed, termbox.ColorDefault, "SPACE TO RESTART")
	}
	printf_String(3*RunConfig.Size+10, 10, "F1:Cancel sign F2:Open F3:Sign Boom F4:Unknow")
	printf_String(3*RunConfig.Size+10, 11, "ESC:EXIT Space:Reset PgUp:Up Size PgDn:Down Size")
}

func printf_String(x, y int, s string) {
	printf_String_Color(x, y, termbox.ColorWhite, termbox.ColorDefault, s)
}
func printf_String_Color(x, y int, fg, bg termbox.Attribute, s string) {
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		w := runewidth.RuneWidth(r)
		if w == 0 || (w == 2 && runewidth.IsAmbiguousWidth(r)) {
			w = 1
		}
		x += w
	}
}
