package snake

import (
	"github.com/nsf/termbox-go"
	"fmt"
	"math/rand"
	"time"
	"errors"
	"strconv"
)

var s [20][20]int

type pos struct{
	x,y int
}
var snake struct{
	d int
	len int
	pos []pos
	keychan chan int
	point int
	start int
	hard int
}

var color = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorYellow,//2
	termbox.ColorMagenta,//4
	termbox.ColorRed,//8
	termbox.ColorGreen,//16
	termbox.ColorCyan,//32
	termbox.ColorBlue,//64
}
const coldef= termbox.ColorDefault
var tk  *time.Ticker

func Snake() {

	i := 0

	sinit()
	random()
	fmt.Println(snake.d,"==")
	for {
		select {
		case <-tk.C:
			ok := snakemove()
			if !ok {
				return
			}
			i++
			ppprint()
		fmt.Println(i)
		default:
		}
	}
}

//0:esc 1:move 2:resatrt
//按键操作
func keyborad(){
	for {
		switch termbox.PollEvent().Key {
		case termbox.KeyEsc:
			fmt.Println("exec game")
			snake.keychan<- 0
		case termbox.KeySpace:
			snake.keychan<- 2

		case termbox.KeyArrowUp:
			if snake.d != 6 {
				snake.d = 4
				//ok := snakemove()
				//if !ok {
				//	fmt.Println("game over")
				//	snake.keychan<- 1
				//	return
				//}
				snake.keychan<- 1
				tker()
			}
			//left 4
		case termbox.KeyArrowDown:
			if snake.d != 4 {
				snake.d = 6
				snake.keychan<- 1
				tker()
			}

			//right 6
		case termbox.KeyArrowRight:
			if snake.d != 2 {
				snake.d = 8
				snake.keychan<- 1
				tker()
			}
			//down 8
		case termbox.KeyArrowLeft:
			if snake.d != 8  {
				snake.d = 2
				snake.keychan<- 1
				tker()
			}
			//up 2
		}
	}
}
//定时器
func timer(){
	for{
		select {
		case <-tk.C:
			snake.keychan <-1
		default:

		}
	}

}
//难度慢慢加大
func tker(){
	if snake.len<10{
		tk = time.NewTicker(750*time.Millisecond)
		snake.hard =1
	}else if snake.len<20{
		tk = time.NewTicker(500*time.Millisecond)
		snake.hard =2
	}else if snake.len<40{
		tk = time.NewTicker(350*time.Millisecond)
		snake.hard =3
	}else if snake.len<80{
		tk = time.NewTicker(275*time.Millisecond)
		snake.hard =4
	}else if snake.len<160{
		tk = time.NewTicker(200*time.Millisecond)
		snake.hard =5
	}else if snake.len<320{
		tk = time.NewTicker(150*time.Millisecond)
		snake.hard =6
	}else {
		tk = time.NewTicker(100*time.Millisecond)
		snake.hard =7
	}
}
//初始化蛇，开始GAME
func SnakeInit() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()
	sinit()

	go keyborad()
	go timer()
	for {
		termbox.Clear(coldef, coldef)
		snakePrintf()
		select {
		case cmd := <-snake.keychan:
			if cmd ==0{
				return
			}else if cmd == 1{
				if snake.start ==0 {
					ok := snakemove()
					if !ok {
						snake.start = 2
					}
				}
			}else if cmd ==2{
				if snake.start ==0{
					snake.start =1
				}else if snake.start==1{
					snake.start=0
				}else {
					snake.start=0
					sinit()
				}

			}
		}
}
}



//随机食物
func random()( error) {
	rand.Seed(time.Now().Unix())
	r := 1
	index := 0
	var a map[int]int
	a = make(map[int]int)
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if s[i][j] != 0 {
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
	for j := 0; j < 20; j++ {
		if s[a[i1]][j] == 0 {
			b[index] = j
			index++
		}
	}
	//fmt.Println(b,len(b))
	i2 := rand.Intn(len(b))
	//fmt.Println(i1,i2,val)
	s[a[i1]][b[i2]]=6
	return nil
}
//蛇移动
func snakemove()bool {
	switch snake.d {
	case 2:
		x, y := snake.pos[0].x, snake.pos[0].y
		//fmt.Println(x,y,s[x][y+1])
		if s[x-1][y] != 0 && s[x-1][y] != 6 && x-1!=snake.pos[snake.len-1].x && y!=snake.pos[snake.len-1].y && snake.len >1{
			return false
		} else if s[x-1][y] != 0 && s[x-1][y] != 6 && snake.len == 1{
			return false
		} else if s[x-1][y] == 6 {
			eat()
		} else {
			move()
		}
	case 8:
		x, y := snake.pos[0].x, snake.pos[0].y
		if s[x+1][y] != 0 && s[x+1][y] != 6 && x+1!=snake.pos[snake.len-1].x && y!=snake.pos[snake.len-1].y && snake.len >1{
			return false
		} else if s[x+1][y] != 0 && s[x+1][y] != 6 && snake.len ==1{
			return false
		}else if s[x+1][y] == 6{
			eat()
		} else {
			move()
		}
	case 4:
		x, y := snake.pos[0].x, snake.pos[0].y
		if s[x][y-1] != 0 && s[x][y-1] != 6 && x!=snake.pos[snake.len-1].x && y-1 !=snake.pos[snake.len-1].y && snake.len >1{
			return false
		}else if s[x][y-1] != 0 && s[x][y-1] != 6 && snake.len == 1{
			return false
		}else if s[x][y-1] == 6 {
			eat()
		} else {
			move()
		}
	case 6:
		x, y := snake.pos[0].x, snake.pos[0].y
		if s[x][y+1] != 0 && s[x][y+1] != 6 && x!=snake.pos[snake.len-1].x && y+1 !=snake.pos[snake.len-1].y && snake.len >1 {
			return false
		} else if  s[x][y+1] != 0 && s[x][y+1] != 6 && snake.len ==1 {
			return false
		}else if s[x][y+1] == 6 {
			eat()
		} else {
			move()
		}

	}
	return true
}
//移动
func move(){
	x,y := snake.pos[0].x,snake.pos[0].y
	switch snake.d {
	case 2:
		snake.pos[0].x--
	case 8:
		snake.pos[0].x++
	case 4:
		snake.pos[0].y--
	case 6:
		snake.pos[0].y++
	}
	//s[snake.pos[0].x][snake.pos[0].y]=3
	for i:=1;i<snake.len;i++{
		x,y ,snake.pos[i].x,snake.pos[i].y =snake.pos[i].x,snake.pos[i].y,x,y
		s[snake.pos[i].x][snake.pos[i].y]=1
	}
	s[x][y]=0
	s[snake.pos[0].x][snake.pos[0].y]=3
}
//吃
func eat(){
	snake.len++
	x,y := snake.pos[0].x,snake.pos[0].y
	switch snake.d {
	case 2:
		snake.pos[0].x--
	case 8:
		snake.pos[0].x++
	case 4:
		snake.pos[0].y--
	case 6:
		snake.pos[0].y++
	}
	s[snake.pos[0].x][snake.pos[0].y]=3
	for i:=1;i<snake.len;i++{
		x,y ,snake.pos[i].x,snake.pos[i].y =snake.pos[i].x,snake.pos[i].y,x,y
		s[snake.pos[i].x][snake.pos[i].y]=1
	}
	snake.point += 10
	random()
}

//初始化
func sinit(){

	tk = time.NewTicker(time.Millisecond*1000)
	snake.pos=make([]pos,400)
	snake.len=1
	snake.d=2
	snake.pos[0].x=9
	snake.pos[0].y=9
	snake.keychan=make(chan int)
	for i:=0;i<20;i++{
		for j:=0;j<20;j++{
			if i==0 || i== 19 || j==0 || j==19{
				s[i][j]=4
			}else {
				s[i][j]=0
			}
		}
	}
	s[snake.pos[0].x][snake.pos[0].y]=3
	random()
}

//绘制
func snakePrintf(){

	for i:=0;i<40;i++{
		for j:=0;j<20;j++{
			termbox.SetCell(i,j,' ',color[s[i/2][j]],color[s[i/2][j]])
		}
	}
	if snake.start == 2 {
		snakePoint(15,10,0,"Game Over")
	}else if snake.start == 1{
		snakePoint(15,10,0,"Game Pause")
	}
	snakePoint(45,4,0,"SCORE:"+strconv.Itoa(snake.point))

	snakePoint(45,8,0,"Hard:"+strconv.Itoa(snake.hard))

	snakePoint(45,6,0,"TIME:"+time.Now().Format("2006-01-02 15:04:05"))

	termbox.Flush()
}

func snakePoint(x,y,z int,buf string ){
	for i:=0;i<len(buf);i++{
		termbox.SetCell(x,y,rune(buf[i]),termbox.ColorYellow,termbox.ColorBlack)
		if z ==0{
			x++
		}else {
			y++
		}
	}
}



func ppprint(){
	fmt.Println("-----------")
	for _,v := range s{
		fmt.Println(v)
	}
}