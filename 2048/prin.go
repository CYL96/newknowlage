package _048

import (
	"github.com/nsf/termbox-go"
	"fmt"
	"math/rand"
	"time"
	"errors"
)
var data [4][4]int

 var color = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorYellow,//2
	termbox.ColorMagenta,//4
	termbox.ColorRed,//8
	termbox.ColorGreen,//16
	termbox.ColorCyan,//32
	termbox.ColorBlue,//64
}

func INIT2048(){
	err := termbox.Init()
	if err !=nil{
		fmt.Println(err)
		return
	}

	random()
	random()
//	Print2048()

	const coldef = termbox.ColorDefault
loop:
	for {
		termbox.Clear(coldef, coldef)

		Print2048()

		switch termbox.PollEvent().Key {
		case termbox.KeyEsc:
			fmt.Println("exec game")
			break loop
		case termbox.KeyArrowUp:
			n := posleft()
			if n ==1{
				err := random()
				if err != nil {
					fmt.Println("game over")
					return
				}
			}

		case termbox.KeyArrowDown:

			n := posright()
			if n ==1{
				err := random()
				if err != nil {
					fmt.Println("game over")
					return
				}
			}
		case termbox.KeyArrowRight:

			n :=posdown()
			if n ==1{
				err := random()
				if err != nil {
					fmt.Println("game over")
					return
				}
			}
		case termbox.KeyArrowLeft:

			n := posup()
			if n ==1{
				err := random()
				if err != nil {
					fmt.Println("game over")
					return
				}
			}
		default:
			goto loop

		}
	//prin()
		//data[0][1]=1
		//data[0][3]=1
		//data[1][3]=2
		//data[1][2]=3
		//data[2][0]=0
		//data[2][1]=2
		//data[2][2]=2
		//data[2][3]=2
		//data[3][0]=1
		//data[3][2]=2
		//data[3][3]=2
		////prin()
		////posw()
		//prin()
		//fmt.Println(6)
		//posss(6)
		//prin()
		//fmt.Println(4)
		//posss(4)
		//prin()
		//fmt.Println(2)
		//posss(2)
		//prin()
		//fmt.Println(8)
		//posss(8)
		//prin()
		//posdown()
		////posright()
		//prin()
		//random()
		//prin()

	}
}


func prin(){
	for _,v := range data{
		fmt.Println(v)
	}
}

//处理 【4】int数据，对齐
func smalldeal(a [4]int,pos int)[4]int{
	t:=0
	for j:=0;j<3;j++{
		if a[t]==0{
			for k:=t;k<3;k++{
				a[k]=a[k+1]
			}
			a[3]=0
		}else {
			t++
		}
	}
	if pos == 2 || pos == 4 {

	}else {
		for j:=0;j<3;j++{
			if a[3]==0{
				a[3],a[2],a[1]=a[2],a[1],a[0]
				a[0]=0
			}

		}
	}
//	fmt.Println(a)
	return a
}

//整体对齐
func posss(pos int) {
	for i:=0;i<4;i++{
		switch pos {
		case 2:
			var a [4]int
			for j:=0;j<4;j++{
				a[j]=data[j][i]
			}
			a = smalldeal(a,2)
			for j:=0;j<4;j++{
				data[j][i]=a[j]
			}
		case 4:
			data[i]=smalldeal(data[i],4)
		case 6:
			data[i]=smalldeal(data[i],6)
		case 8:
			var a [4]int
			for j:=0;j<4;j++{
				a[j]=data[j][i]
			}
			a = smalldeal(a,8)
			for j:=0;j<4;j++{
				data[j][i]=a[j]
			}

		}
	}

}

//左移
func posleft()int {
	a := data
	posss(4)

	for i:=0;i<4;i++{
		for j:=0;j<3;j++{
			if data[i][j] == data[i][j+1] && data[i][j] != 0 && data[i][j] <6 {
				data[i][j] +=1
				data[i][j+1]=0
			}
		}
	}
	posss(4)
	if a == data{
		return 0
	}
	return 1
}
//右移
func posright()int {
	a := data
	posss(6)

	for i:=0;i<4;i++{
		for j:=3;j>0;j--{
			if data[i][j] == data[i][j-1] && data[i][j] != 0 && data[i][j] <6 {
				data[i][j] +=1
				data[i][j-1]=0
			}
		}
	}
	posss(6)
	if a == data{
		return 0
	}
	return 1
}
func posup()int{
	a := data
	posss(2)
	for i:=0;i<4;i++{
		for j:=0;j<3;j++{
			if data[j][i]==data[j+1][i] && data[j][i] !=0 && data[j][i]<6{
				data[j][i] += 1
				data[j+1][i]=0
			}
		}
	}
	posss(2)
	if a == data{
		return 0
	}
	return 1
}
func posdown()int{
	a:= data
	posss(8)

	for i:=0;i<4;i++{
		for j:=3;j>0;j--{
			if data[j][i]==data[j-1][i] && data[j][i] !=0 && data[j][i]<6{
				data[j][i] += 1
				data[j-1][i]=0
			}
		}
	}
	posss(8)
	if a == data{
		return 0
	}
	return 1
}
//随机种子
func random()( error) {
	rand.Seed(time.Now().Unix())
	r := 1
	index := 0
	var a map[int]int
	a = make(map[int]int)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if data[i][j] != 0 {
				r++
			}
		}
		if r < 4 {
			a[index] = i
			index++
		}
		r=0
	}
	if len(a) == 0 {
		return errors.New("game over")
	}
	i1 := rand.Intn(len(a))
	//fmt.Println(a,len(a))
	var b map[int]int
	b = make(map[int]int)
	index=0
	for j := 0; j < 4; j++ {
		if data[a[i1]][j] == 0 {
			b[index] = j
			index++
		}
	}
	//fmt.Println(b,len(b))
	i2 := rand.Intn(len(b))
	val := rand.Intn(2)+1
	//fmt.Println(i1,i2,val)
	data[a[i1]][b[i2]]=val
	return nil
}


func Print2048(){
	for i:= 0;i<41;i++{
		for j:=0;j<25;j++{
			if i%10==0{
				termbox.SetCell(i,j,'|',termbox.ColorRed,termbox.ColorWhite)
			}else if j%6==0 {
				termbox.SetCell(i,j,'-',termbox.ColorRed,termbox.ColorWhite)
			}else {
				//fmt.Println(data[i/10][j/6])
				termbox.SetCell(i,j,' ',color[data[i/10][j/6]],color[data[i/10][j/6]])
			}
			//else if i%5==0 && j%3==0{
			//	termbox.SetCell(i,j,' ',termbox.ColorRed,termbox.ColorWhite)
			//}
		}
	}
	for i := 50;i<57;i++{
		for j:=0;j<27;j++ {
				//fmt.Println(data[i/10][j/6])
				termbox.SetCell(i, j, ' ', color[j/4], color[j/4])
		}
	}
	termbox.Flush()
}