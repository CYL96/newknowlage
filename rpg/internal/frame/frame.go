package frame

import (
	"fmt"
	"os"

	"newknowlage/rpg/internal/common"
	. "newknowlage/rpg/internal/gui"
)

func Run() {
	var (
		err error
	)
	err = os.Setenv("FYNE_FONT", "./config/wryh.ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Unsetenv("FYNE_FONT")

	err = common.InitUrlConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = InitGui()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	GetMyGui().Run()

}
