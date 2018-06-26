package ch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

func CHtest1() {

	//wait := Publish("1111111",4*time.Second)
	//fmt.Println(111)
	//fmt.Println(<-wait)

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go race(&wg)
	}

	wg.Wait()

}
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	fmt.Println(1)
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) // 广播 - 一个关闭的管道都会发送一个零值
	}()
	fmt.Print(1)
	return ch
}

func race(wg *sync.WaitGroup) {
	wait := make(chan struct{})
	n := 0
	go func() {
		// 译注：注意下面这一行
		n++ // 一次访问: 读, 递增, 写
		close(wait)
		wg.Done()
	}()
	// 译注：注意下面这一行
	n++ // 另一次冲突的访问
	<-wait
	fmt.Print(n, "=") // 输出：未指定

}

func SendmessageToWX(user string, msg string) {
	url := "https://pm.cninct.com/SendMessage"

	buf := `type=WX&code=10000&sign=6cfbc2d63ad294d7d3b2360b606afb71&user=["` + user + `"]&Message=` + msg

	payload := strings.NewReader(buf)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func QQQQQQ(unit string) {
	if unit == "小时" {

	} else if unit == "天" {

	} else if unit == "周" {

	}
}
