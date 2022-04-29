package test

import "fmt"

type Number interface {
	int64 | float64 |int | float32
}

func GOOGOGO(){
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
	fmt.Println(SumNumbers(ints))
	fmt.Println(ReturnNumber2(1,2),ReturnNumber2(1.11,2.0),Pkgv(1))
}

func ReturnNumber2[V int64|float64|int|int32|float32](n V,b V) V {
	return n+b
}
func ReturnNumber[V Number](n V,b V) V {
	return n+b
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}