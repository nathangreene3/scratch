package scratch

import "fmt"

func toInt(x interface{}) int {
	switch t := x.(type) {
	case int:
		return x.(int)
	case *int:
		return int(*x.(*int))
	default:
		panic(fmt.Sprintf("unsupported type %t", t))
	}
}

func sum(values []int) int {
	var s int
	n := len(values)
	for i := 0; i < n; i++ {
		s += values[i]
	}
	return s
}

func sum0(values []int) int {
	var s int
	for i := range values {
		s += values[i]
	}
	return s
}

func sum1(values []int) int {
	var s int
	for _, v := range values {
		s += v
	}
	return s
}

func sum2(values ...int) int {
	var s int
	for i := range values {
		s += values[i]
	}
	return s
}

func sum3(values ...int) int {
	var s int
	for _, v := range values {
		s += v
	}
	return s
}

func sum4(values []int) int {
	return sumr0(0, len(values), values)
}

func sumr0(i, n int, values []int) int {
	if i < n {
		return values[i] + sumr0(i+1, n, values)
	}
	return 0
}

func sum5(values ...int) int {
	return sumr1(0, len(values), values...)
}

func sumr1(i, n int, values ...int) int {
	if i < n {
		return values[i] + sumr1(i+1, n, values...)
	}
	return 0
}
