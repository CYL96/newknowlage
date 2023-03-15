package main

import "fmt"

//  -------------------------------------- 转换为int64 ------------------------------------------------
/*
//  AnyIntToInt64[T int | int8 | int16 | int32 | int64]
//  @Description:
//  @param [v]: 需要转换的值
//  @return [int64]: 返回的int64
*/
func AnyIntToInt64[T int | int8 | int16 | int32 | int64](v T) int64 {
	return int64(v)
}

//  -------------------------------------- 返回最大的值 ------------------------------------------------
/*
//  WhoIsBig[T int | int8 | int16 | int32 | int64 | float32 | float64]
//  @Description:
//  @param [a]: 值1
//  @param [b]: 值2
//  @return [T]: 最大的值
*/
func WhoIsBig[T int | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

//  -------------------------------------- 三元计算 ------------------------------------------------
/*
//  TernaryCalculation[T any]
//  @Description:
//  @param [f]: 比较的语句
//  @param [a]: a
//  @param [b]:  b
//  @return [T]: f = true 时返回a  f = false 时返回 b
*/
func TernaryCalculation[T any](f bool, a, b T) T {
	if f {
		return a
	}
	return b
}

func main() {
	a := 111111
	fmt.Println(AnyIntToInt64(a))
	fmt.Println(WhoIsBig(100, 111))
	fmt.Println(WhoIsBig(101.231, 111.1))

	fmt.Println(TernaryCalculation(100 > 102, "成功了", "失败了"))
	fmt.Println(TernaryCalculation(100 < 102, "成功了", "失败了"))
}
