package core

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"unicode/utf8"

	. "gosdk/qrcode"
)

type drawWordContentT struct {
	content  string
	num      int
	fontSize int
	fontW    int
	lineSize int
	BackSize int `xml:"back_size"`
	point    drawWordPoint
}
type drawWordPoint struct {
	X int
	Y int
}
type drawWordT struct {
	totalHigh int
	word      []drawWordContentT
}

func Drawwww() {
	pic, err := NewDrawPic("./config/welcome.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pic.InitFront("./config/welcome.ttf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pic.SetFontDpi(70)
	var data drawWordT
	for i, s := range Config.Info.Desc {
		tmp := drawWordContentT{
			content:  s.Content,
			num:      utf8.RuneCountInString(s.Content),
			fontSize: s.Size,
			fontW:    s.Kerning,
			lineSize: s.Line,
			BackSize: s.BackSize,
		}
		data.word = append(data.word, tmp)
		if i != 0 {
			data.totalHigh += tmp.lineSize
		}
		data.totalHigh += tmp.fontSize
	}
	lastY := 0

	for i, t := range data.word {
		t.point.X = (3840 - t.num*t.fontSize - t.num*t.fontW) / 2
		if lastY == 0 {
			t.point.Y = (2160-data.totalHigh)/2 + int(float64(t.fontSize)*0.7777)
			lastY = t.point.Y
		} else {
			t.point.Y = lastY + t.lineSize + t.fontSize
			lastY = t.point.Y
		}

		line := []DrawFontBatchT{}
		content := []DrawFontBatchT{}
		lineNum := t.BackSize

		for _, i3 := range t.content {
			line = append(line, wordMoveTo(string(i3), t.point.X-lineNum, t.point.Y-lineNum, t.point.X+lineNum, t.point.Y+lineNum)...)
			content = append(content, DrawFontBatchT{
				Content: string(i3),
				X:       t.point.X,
				Y:       t.point.Y,
			})
			// pic.DrawFont(string(i3), t.point.X, t.point.Y)
			t.point.X += t.fontSize + t.fontW
		}
		pic.SetFontSize(float64(t.fontSize))
		// 设置轮廓
		pic.SetFontColor(Config.Color.Back.R, Config.Color.Back.G, Config.Color.Back.B, Config.Color.Back.A)
		fmt.Println("写入背景中：", t.content)
		pic.DrawFontBatch(line)

		fmt.Println("写入内容：", t.content)
		if i == 0 || i == len(data.word)-1 {
			// pic.SetFontColor(245, 245, 0xff, 255)
			pic.SetFontColor(Config.Color.Title.R, Config.Color.Title.G, Config.Color.Title.B, Config.Color.Title.A)
			// pic.SetFontColor(27, 91, 97, 0xff)
		} else {
			// pic.SetFontColor(0xff, 241, 0, 0xff)
			pic.SetFontColor(Config.Color.Content.R, Config.Color.Content.G, Config.Color.Content.B, Config.Color.Content.A)
		}
		pic.DrawFontBatch(content)

	}
	err = pic.Save("./welcome.png", 100, PIC_TYPE_PNG)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	tC := color.RGBA{Config.Color.Title.R, Config.Color.Title.G, Config.Color.Title.B, Config.Color.Title.A}
	cC := color.RGBA{Config.Color.Content.R, Config.Color.Content.G, Config.Color.Content.B, Config.Color.Content.A}
	// bC := color.RGBA{Config.Color.Back.R, Config.Color.Back.G, Config.Color.Back.B, Config.Color.Back.A}
	newPic := pic.DefaultPic.(*image.NRGBA)
	max_x := pic.DefaultPic.Bounds().Size().X
	max_y := pic.DefaultPic.Bounds().Size().Y
	// draw.Draw(newPic, newPic.Bounds(), pic.DefaultPic, image.ZP, draw.Src)
	width := 50
	for x := 0; x < max_x; x++ {
		for y := 0; y < max_y; y++ {
			bC := color.RGBA{Config.Color.Back.R, Config.Color.Back.G, Config.Color.Back.B, Config.Color.Back.A}
			color_o := pic.DefaultPic.At(x, y)
			// newPic.Set(x, y, pic.DefaultPic.At(x, y))
			if x == 0 || x >= max_x-width || y == 0 || y >= max_y-width {
				continue
			}
			if rgbaC(color_o, tC) || rgbaC(color_o, cC) && color_o != bC {
				for i := 1; i <= width; i++ {
					if width > 2 && (i == 1) {
						// || i == width
						rgba := color_o.(color.NRGBA)
						bC = color.RGBA{colorTransition(Config.Color.Back.R, rgba.R),
							colorTransition(Config.Color.Back.G, rgba.G),
							colorTransition(Config.Color.Back.B, rgba.B),
							colorTransition(Config.Color.Back.A, rgba.A)}
					} else {
						bC = color.RGBA{Config.Color.Back.R, Config.Color.Back.G, Config.Color.Back.B, Config.Color.Back.A}
					}
					color_u := pic.DefaultPic.At(x, y+i)
					color_d := pic.DefaultPic.At(x, y-i)
					color_l := pic.DefaultPic.At(x-i, y)
					color_r := pic.DefaultPic.At(x+i, y)
					if !rgbaC(color_u, tC) && !rgbaC(color_u, cC) {
						newPic.Set(x, y+i, bC)
					}
					if !rgbaC(color_d, tC) && !rgbaC(color_d, cC) {
						newPic.Set(x, y-i, bC)
					}
					if !rgbaC(color_l, tC) && !rgbaC(color_l, cC) {
						newPic.Set(x-i, y, bC)
					}
					if !rgbaC(color_r, tC) && !rgbaC(color_r, cC) {
						newPic.Set(x+i, y, bC)
					}
				}
			}
		}
	}

	imgF, err := os.Create("./welcome_1.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = png.Encode(imgF, newPic)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("完成")

}

func rgbaC(l, f color.Color) bool {
	l_r, l_g, l_b, l_a := l.RGBA()
	f_r, f_g, f_b, f_a := f.RGBA()
	if l_r == f_r && l_g == f_g && l_b == f_b && l_a == f_a {
		return true
	}
	return false
}

func colorTransition(l, r uint8) uint8 {
	if l > r {
		return r + (l - r)
	} else {
		return l + (r - l)
	}
}
func wordMoveTo(content string, x, y, xTo, yTo int) (list []DrawFontBatchT) {
	for i := 0; i < xTo-x; i++ {
		list = append(list, DrawFontBatchT{
			Content: content,
			X:       x + i,
			Y:       y,
		})
		list = append(list, DrawFontBatchT{
			Content: content,
			X:       x + i,
			Y:       yTo,
		})
	}
	for i := 0; i < yTo-y; i++ {
		list = append(list, DrawFontBatchT{
			Content: content,
			X:       x,
			Y:       y + i,
		})
		list = append(list, DrawFontBatchT{
			Content: content,
			X:       xTo,
			Y:       y + i,
		})
	}
	return
}
