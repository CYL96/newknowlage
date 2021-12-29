package boom

import (
	"github.com/nsf/termbox-go"
	"os"
)

var RunConfig Struct_Config

type Struct_Config struct {
	Area      [][]show
	Size      int
	MineNum   int
	IsMine    int
	TrueMine  int
	Opened    int
	Start     int
	Key       chan termbox.Key
	UseTime   int64
	StartTime int64
	log       *os.File
	x         int
	y         int
}

const (
	reset  = 0
	open   = 1
	isMine = 2
	unknow = 3
)

type show struct {
	Open    int
	MineNum int32
	Mine    bool
}
