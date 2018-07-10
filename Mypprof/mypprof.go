package Mypprof

import (
	"os"
	"runtime/pprof"
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
