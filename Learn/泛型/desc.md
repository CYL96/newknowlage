>golang 的泛型是在 1.18 版本中才正式加入的，但是在 1.17 版本中已经可以使用了，只是需要加上 -gcflags=-G=3 参数才可以使用。

## 1. 泛型的使用
示例1：
```golang
func AnyIntToInt64[T int | int8 | int16 | int32 | int64](v T) int64 {
	return int64(v)
}
```
示例2：
```golang
func WhoIsBig[T int | int8 | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```
示例3:
```golang
func TernaryCalculation[T any](f bool, a, b T) T {
	if f {
		return a
	}
	return b
}
```