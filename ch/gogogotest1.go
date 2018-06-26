package ch

import "fmt"

func GOGO() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
