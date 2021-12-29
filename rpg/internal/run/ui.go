package run

import "fmt"

func Page0() {
	User.Rebuild()
	Monster.Rebuild()

	User.ShowStatus()
	Monster.ShowStatus()
	fmt.Println("1: attack 2:exit")
	nowPage = GetWantInput(1, 2)
}

var nowPage int
var exit int

const (
	firstPage_num    = 0
	search_num       = 3
	searchResult_num = 4
)

func Go() {
	for {
		Clear()
		switch nowPage {
		case 0:
			Page0()
			if exit == 1 {
				return
			}
		case 1:
			ok := Attack()
			if ok {
				User.Upgrade()
				nowPage = 1
			} else {
				nowPage = 0
			}
		case 2:
			return
		}
	}

}
