package boom

import (
	"fmt"
	"math/rand"
	"time"
)

func SetSizeAndNum(size int) {
	if size < 5 {
		size = 5
	}
	RunConfig.Size = size
}

func SetMine() {
	if RunConfig.Size < 5 {
		RunConfig.Size = 5
	}
	RunConfig.MineNum = (RunConfig.Size-5)*4 + 5
	RunConfig.x, RunConfig.y = 0, 0

	RunConfig.Area = make([][]show, RunConfig.Size)
	for i := range RunConfig.Area {
		RunConfig.Area[i] = make([]show, RunConfig.Size)
	}

	data := make(map[string]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < RunConfig.MineNum; i++ {
	MINE:
		x := rand.Intn(RunConfig.Size)
		y := rand.Intn(RunConfig.Size)
		if _, ok := data[fmt.Sprintf("%d,%d", x, y)]; ok {
			goto MINE
		} else {
			RunConfig.Area[y][x].Mine = true
			data[fmt.Sprintf("%d,%d", x, y)] = 1
		}
	}
}
func SetMineNum() {
	for y := range RunConfig.Area {
		for x := range RunConfig.Area[y] {
			if RunConfig.Area[y][x].Mine {
				continue
			}
			RunConfig.Area[y][x].MineNum = getMineNum(x, y)
		}
	}
}

func getMineNum(x, y int) int32 {
	var min_x, max_x, min_y, max_y int
	var MineNum int32
	if y == 0 {
		min_y = 0
		max_y = 1
	} else if y == RunConfig.Size-1 {
		min_y = RunConfig.Size - 2
		max_y = RunConfig.Size - 1
	} else {
		min_y = y - 1
		max_y = y + 1
	}
	if x == 0 {
		min_x = 0
		max_x = 1
	} else if x == RunConfig.Size-1 {
		min_x = RunConfig.Size - 2
		max_x = RunConfig.Size - 1
	} else {
		min_x = x - 1
		max_x = x + 1
	}
	//fmt.Println(y,x,min_x,max_x,min_y,max_y)
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			if RunConfig.Area[y][x].Mine {
				MineNum++
			}
		}
	}
	return 48 + MineNum
}

func SetFlag(x, y, action int) bool {
	if RunConfig.Start == 0 {
		switch action {
		case reset:
			if RunConfig.Area[y][x].Open != open {
				if RunConfig.Area[y][x].Open == isMine {
					RunConfig.IsMine--
				}

				RunConfig.Area[y][x].Open = reset
			}
		case open:
			if RunConfig.Area[y][x].Open != open {
				if RunConfig.Area[y][x].Open == isMine {
					RunConfig.IsMine--
					if RunConfig.Area[y][x].Mine {
						RunConfig.TrueMine--
					}
				}
				RunConfig.Area[y][x].Open = open
				RunConfig.Opened++
				if RunConfig.Area[y][x].Mine {
					return false
				}
				openMineNum(x, y)

			}

		case isMine:
			if RunConfig.Area[y][x].Open != open {
				if RunConfig.IsMine != RunConfig.MineNum {
					RunConfig.IsMine++
					RunConfig.Area[y][x].Open = isMine
					if RunConfig.Area[y][x].Mine {
						RunConfig.TrueMine++
					}
				}
			}
		case unknow:
			if RunConfig.Area[y][x].Open != open {

				if RunConfig.Area[y][x].Open == isMine {
					RunConfig.IsMine--
					if RunConfig.Area[y][x].Mine {
						RunConfig.TrueMine--
					}
				}

				RunConfig.Area[y][x].Open = unknow

			}

		}
	}
	return true
}
func openMineNum(x, y int) {
	if RunConfig.Area[y][x].MineNum-48 != 0 {
		//fmt.Println(RunConfig.Area[y][x].MineNum)
		return
	}
	var min_x, max_x, min_y, max_y int
	if y == 0 {
		min_y = 0
		max_y = 1
	} else if y == RunConfig.Size-1 {
		min_y = RunConfig.Size - 2
		max_y = RunConfig.Size - 1
	} else {
		min_y = y - 1
		max_y = y + 1
	}
	if x == 0 {
		min_x = 0
		max_x = 1
	} else if x == RunConfig.Size-1 {
		min_x = RunConfig.Size - 2
		max_x = RunConfig.Size - 1
	} else {
		min_x = x - 1
		max_x = x + 1
	}
	//fmt.Println(y,x,min_x,max_x,min_y,max_y)
	//RunConfig.log.WriteString(fmt.Sprintf("1: %d,%d,%d,%d,%d,%d \n",y,x,min_x,max_x,min_y,max_y))
	for y_t := min_y; y_t <= max_y; y_t++ {
		for x_t := min_x; x_t <= max_x; x_t++ {
			if x_t == x && y_t == y {
				continue
			}
			//RunConfig.log.WriteString(fmt.Sprintf("%d,%d,%d,%d \n",y,x,y_t,x_t))
			if RunConfig.Area[y_t][x_t].Open == reset && !RunConfig.Area[y_t][x_t].Mine {
				RunConfig.Area[y_t][x_t].Open = open
				RunConfig.Opened++
				openMineNum(x_t, y_t)
			}
		}
	}

	return
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func GameOver() {
	if RunConfig.Start == 1 {
		return
	}
	RunConfig.UseTime = time.Now().Unix() - RunConfig.StartTime
	for y := range RunConfig.Area {
		for x := range RunConfig.Area[y] {
			if RunConfig.Area[y][x].Mine {
				RunConfig.Area[y][x].Open = 1
			}
		}
	}
}
