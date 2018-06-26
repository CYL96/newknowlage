package download

import (
	"net/http"
	"fmt"
	"strings"
	"sync"
	"time"
	"strconv"
	"os"
	"io"
)

var filebuf struct{
	lk *sync.RWMutex
	buf []byte
	dlsize int64
	totalsize int64
}
var speed struct{
	starttime int64
	endtime int64
	average float32
	realtime float32
}

func timer(){
	tk := time.NewTicker(1*time.Second)
	oldsize := int64(0)
	for{
		select {
		case <-tk.C:
			if filebuf.dlsize<filebuf.totalsize{
				filebuf.lk.Lock()
				size := filebuf.dlsize-oldsize
				oldsize =filebuf.dlsize
				speed.realtime= float32(size)/1024
				fmt.Printf("Total:%.2f KB;Dowload:%.2f KB;Speed:%.2f KB/S \r",float32(filebuf.totalsize)/1024,float32(filebuf.dlsize)/1024,speed.realtime)
				filebuf.lk.Unlock()
			}

		}
	}
}
func DownLoadTest(){
	url := `http://f8.volumes.cc/08e8cf0d1530003111/0082/50791/%5BMobi%5D%5BVol.moe%5D%5Bjjs%5DVol_003.mobi`
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url)
	file := strings.SplitAfter(url,"/")
	fmt.Println("file:",file[len(file)-1],"start download,please wait.")
	for i,v := range resp.Header{
		fmt.Println(i,":",v)
	}
}





func DownLoadInit(url string)error{
	//url := "http://dl.pandownload.com/download/PanDownload_v1.5.4.zip"
	//url = `http://mydmplus.com/res/Mydm/Mydm20180623.zip`
	//url := `http://dl186.80s.im:920/1805/头号玩家/头号玩家.mp4`

	filebuf.lk = new(sync.RWMutex)
	filebuf.buf = make([]byte,10*4096)

	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	filebuf.totalsize,err = strconv.ParseInt( resp.Header.Get("Content-Length"),10,64)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(filebuf.totalsize)
	file := strings.SplitAfter(url,"/")
	fmt.Println("file:",file[len(file)-1],"start download,please wait.")
	err = pathcreate("./file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	go timer()
	fd,err:= os.Create("./file/"+file[len(file)-1])
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i,v := range resp.Header{
		fmt.Println(i,":",v)
	}

	speed.starttime=time.Now().Unix()
	for{
		filebuf.lk.Lock()
		n ,err := io.ReadFull(resp.Body,filebuf.buf)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			if err.Error() == "unexpected EOF"{
				filebuf.dlsize += int64(n)
				fd.Write(filebuf.buf[:n])
				speed.endtime=time.Now().Unix()
				fmt.Println()
				fmt.Printf("Total:%.2f KB;Dowload:%.2f KB;Speed:%.2f KB/S \n",float32(filebuf.totalsize)/1024,float32(filebuf.dlsize)/1024,speed.realtime)
				fmt.Println(file[len(file)-1],"download successed!Save to ./file/")
				totalsecond := speed.endtime-speed.starttime
				speed.average = float32((filebuf.totalsize/1024)/totalsecond)
				fmt.Println("Average speed:",speed.average," KB/S")
				filebuf.lk.Unlock()
				return  nil
			}else {
				filebuf.lk.Unlock()
				fmt.Println(11)
				return err
			}

		}
		filebuf.dlsize += int64(n)
		fd.Write(filebuf.buf)
		filebuf.lk.Unlock()
	}
	return nil
}

//自动创建文件夹
func pathcreate(path string) error {
	var err error
	_, err = os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.MkdirAll(path, 0711)
	}
	return err
}

func downloadyes(){

}