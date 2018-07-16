package Mypprof

import (
	"os"
	"runtime/pprof"
	"time"
	"fmt"
)

func Mypprofstart(){
	f, err := os.Create("cpuprofile.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
}
func MypprofStop(){
	pprof.StopCPUProfile()
}

// 生成 CPU 报告
func MyCpuProfile(t time.Duration) {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	if t <=30*time.Second{
		time.Sleep(30 * time.Second)
	}else {
		time.Sleep(t)
	}

	fmt.Println("CPU Profile stopped")
}

// 生成堆内存报告
func MyHeapProfile(t time.Duration) {
	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if t <=30{
		time.Sleep(30 * time.Second)
	}else {
		time.Sleep(t)
	}
	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}
