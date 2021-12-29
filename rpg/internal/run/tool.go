package run

import (
	"fmt"
	"strconv"
)

func GetWantInput(rang ...int) (res int) {
	a := 0
	for {
		fmt.Scan(&a)
		if len(rang) == 0 {
			return -1
		}
		for _, v := range rang {
			if a == v {
				return a
			}
		}
		fmt.Println("输入错误！请重新输入：")
	}
}
func Decimal(value float64, fix int) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(fix)+"f", value), 64)
	return value
}
