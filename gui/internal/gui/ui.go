package gui

func RunUI(uiIndex int) {
	switch uiIndex {
	case 1:
		// 	主页面
		MainUI()
	case 2:
		GetNewList(1, 1, "")
	case 3:
		GetNewList(1, 2, "")
	case 4:
		GetNewList(1, 3, "")
	}
}
