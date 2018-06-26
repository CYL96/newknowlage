package sort

import "fmt"

//快速排序
func Qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	fmt.Println(data)

	Qsort(data[:head])
	Qsort(data[head+1:])
}

//冒泡排序
func MPsort(a []int){
	n := len(a)
	fmt.Println("初始：",a)
	for i := 0;i < n-1;i++{
		for j := 0;j < n-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
		fmt.Println("第",i,"次：",a)
	}
}

//归并排序
func GBsort(a []int) []int{
	n := len(a)
	fmt.Println(n,"--",a)
	if n <=1{
		return a
	}
	lsize := len(a[:n/2])
	left := a[:n/2]
	if lsize >=2 {
		left = GBsort(a[:n/2])
	}
	rsize := len(a[n/2:])
	right := a[n/2:]
	if rsize >=2 {
		right = GBsort(a[n/2:])
	}
	l,r := 0,0
	tmp := make([]int,n)
	for i := 0;i <n;i++{
		if l<lsize && r <rsize{
			if left[l] <= right[r]{
				tmp[i] = left[l]
				l++
			}else  {
				tmp[i] = right[r]
				r++
			}
		}else if l >= lsize{
			tmp[i] = right[r]
			r++
		}else if r >=rsize{
			tmp[i] = left[l]
			l++
		}
	}
	fmt.Println(n,"==",tmp)
	return tmp

}

//选择排序
func XZsort(a []int) {
	n := len(a)
	var min int
	for i :=0;i<n;i++{
		min =i
		for j := i+1 ;j<n;j++{
			if a[j]< a[min]{
				min =j
			}
		}
		a[i],a[min]=a[min],a[i]
	}
}
