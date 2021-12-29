package gui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	. "newknowlage/gui/internal/common"
	. "newknowlage/gui/internal/core"
)

var RunPath string

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	RunPath = dir
	// RunPath = strings.Replace(dir, "\\", "/", -1)
}
func SetUrlLabel() {
	var (
		con *fyne.Container
	)
	gui := GetMyGui()
	con = container.NewVBox()
	urlConfig := GetUrlConfig()
	for _, ext := range urlConfig.UrlList {
		con.Add(widget.NewHyperlink(ext.Name, ext.UrlS))
	}
	bu := widget.NewButton("Hi!", func() {
		Button()
	})
	con.Add(bu)
	gui.MainWindow.SetContent(con)
}

func MainUI() {
	var (
		con *fyne.Container
	)
	gui := GetMyGui()
	gui.MainWindow.Resize(fyne.Size{
		Width:  1000,
		Height: 500,
	})
	gui.MainWindow.SetFixedSize(true)
	con = container.NewVBox()
	bu := widget.NewButton("获取当前最新", func() {
		RunUI(2)
	})
	con.Add(bu)
	bu2 := widget.NewButton("获取当前最多下载", func() {
		RunUI(3)
	})
	con.Add(bu2)
	bu3 := widget.NewButton("搜索", func() {
		RunUI(4)
	})
	con.Add(bu3)
	gui.MainWindow.SetContent(con)
}

type GtButton struct {
	*widget.Button
	Name string
	Id   string
	gt   *GtButton
}

func GetNewList(page int, fg int, search string) {
	var (
		gtconV1, gtconV2, gtconH *fyne.Container
		swH                      *fyne.Container
		mainV                    *fyne.Container
		data                     []SearchResultExt
		err                      error
	)
	mainV = container.NewVBox()
	gui := GetMyGui()
	gtconV1 = container.NewVBox()
	gtconV2 = container.NewVBox()
	gtconH = container.NewHBox()

	if page <= 0 {
		page = 1
	}
	switch fg {
	case 1:
		// 	最新的
		data, err = GetNewGT(page)
	case 2:
		// 	最多下载
		data, err = GetHotGT(page)
	case 3:
		// 	搜索

		sss := widget.NewEntry()
		sss.SetText(search)
		mainV.Add(sss)

		sbut := widget.NewButton("搜索", func() {
			search = sss.Text
			GetNewList(page, fg, search)
		})
		mainV.Add(sbut)

		data, err = GetSearchResult(search, page)
		// form Newfo
	}
	if err != nil {
		lab := widget.NewLabel(err.Error())
		gtconH.Add(lab)
	} else {
		if len(data) == 0 {
			lab := widget.NewLabel("暂无搜索内容")

			// lab.Resize(fyne.NewSize(200, 50))
			mainV.Add(lab)
		}
		for i, datum := range data {
			var (
				bugt *widget.Button
			)
			bugt = func(dt SearchResultExt) *widget.Button {
				return widget.NewButton(dt.SimpleName, func() {
					bugt.Disable()
					bugt.SetText(dt.SimpleName + "-下载中....")
					err = DownGTPic("GT/"+dt.Name, dt.Id)
					if err != nil {
						bugt.SetText(dt.SimpleName + "-下载失败：" + err.Error())
					} else {
						bugt.SetText(dt.SimpleName + "-下载成功")
						if len(RunPath) != 0 && runtime.GOOS == "windows" {
							exec.Command(`cmd`, `/c`, `explorer`, RunPath+"\\"+"GT\\"+dt.Name).CombinedOutput()
						}
					}
				})
			}(datum)
			if i%2 == 0 {
				gtconV1.Add(bugt)
			} else {
				gtconV2.Add(bugt)
			}
		}
	}
	gtconH.Add(gtconV1)
	gtconH.Add(gtconV2)
	mainV.Add(gtconH)
	// 添加下部分操作按钮
	swH = container.NewHBox()
	bu := widget.NewButton("上一页", func() {
		GetNewList(page-1, fg, search)
	})
	if page == 1 {
		bu.Disable()
	}
	swH.Add(bu)
	lab := widget.NewLabel(fmt.Sprintf("第%d页", page))
	swH.Add(lab)

	bu2 := widget.NewButton("下一页", func() {
		GetNewList(page+1, fg, search)
	})
	swH.Add(bu2)
	if len(data) < 25 {
		bu2.Disable()
	}

	bu3 := widget.NewButton("返回", func() {
		RunUI(1)
	})
	swH.Add(bu3)
	mainV.Add(swH)

	gui.MainWindow.SetContent(mainV)
	gui.MainWindow.Content().Refresh()
	gui.MainWindow.Resize(mainV.MinSize())
}
func Button() {
	var (
		con *fyne.Container
	)
	gui := GetMyGui()
	con = container.NewVBox()
	bu := widget.NewButton("Hi!", func() {
		SetUrlLabel()
	})
	con.Add(bu)
	gui.MainWindow.SetContent(con)
}
