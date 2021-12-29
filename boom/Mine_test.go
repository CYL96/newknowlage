package boom

import (
	"fmt"
	"testing"
)

func TestPrintfMine(t *testing.T) {
	SetSizeAndNum(5)
	SetMine()
	SetMineNum()

	PrintfMine()
	var x, y, a int
	for {
		fmt.Scan(&x, &y, &a)
		x--
		y--
		ok := SetFlag(x, y, a)
		PrintfMine()
		if !ok {
			return
		}
	}
}
