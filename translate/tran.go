package translate

import (
	"bufio"
	"fmt"
	"os"
//	_"github.com/mattn/go-sqlite3"
	"io"
)


var fi *os.File
var search []*bufio.Reader //= make([]*bufio.Reader,4,10)
var data chan string

func TranslateInit(path ...string) {
	opentranslatefile(path...)
	data = make(chan string)
}

func TranslateFromTxt(word string) string{
	if len(search) == 0{
		return "NO FILE"
	}
	for _,v := range search{
		go deal(word,v,data)
	}
	var tdata string
	for i := len(search);i>0;i--{
		tdata += <- data
		tdata +="\n"
	}
	return tdata
}


func opentranslatefile(path ...string) (error) {
	for _,v := range path{
		fi, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
			return err
		}
		search = append(search,bufio.NewReader(fi))
	}

	return  nil
}
func deal(word string,sea *bufio.Reader,ch chan string){
	word +=" "
	for {
		line, _, err := sea.ReadLine()
		if err == io.EOF {
			ch <- ""
			return
		}
		if string(line[:len(word)])== word{
			i:=len(word)
			for {
				if string(line[i]) != " "{
					break
				}
				i++
			}
			ch <- string(line[i:])
			//			fmt.Println(ttt,len(ttt),string(line[i:]),string(line[:]))
			return
		}
	}
	return
}



func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}
	return string(r[start : end])
}

//func TxtTOsql(path string) {
//	db, err := sql.Open("")
//}
