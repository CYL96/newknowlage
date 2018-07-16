package test

import "fmt"

func MyMap(){
	fmt.Println("MyMap:")
	buf := make(map[string]string)
	buf["a"]="aa"
	buf["b"]="bb"
	buf["c"]="cc"
	buf["d"]="dd"
	buf["e"]="ee"
	buf["1111"]="111"
	for i,v :=range buf{
		fmt.Println(i,v)
	}
}
