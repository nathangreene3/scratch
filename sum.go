package scratch

func sum0(values []int) int {
	var (
		s int
		n = len(values)
	)

	for i := 0; i < n; i++ {
		s += values[i]
	}

	return s
}

func sum1(values []int) int {
	var s int
	for i := range values {
		s += values[i]
	}

	return s
}

func sum2(values []int) int {
	var s int
	for _, v := range values {
		s += v
	}

	return s
}

func sum3(values ...int) int {
	var (
		s int
		n = len(values)
	)

	for i := 0; i < n; i++ {
		s += values[i]
	}

	return s
}

func sum4(values ...int) int {
	var s int
	for i := range values {
		s += values[i]
	}

	return s
}

func sum5(values ...int) int {
	var s int
	for _, v := range values {
		s += v
	}

	return s
}

// sum6 adds values.
func sum6(values []int) int {
	return sumr6(0, len(values), values)
}

// sumr6 recursively adds values over the range [a,b).
func sumr6(a, b int, values []int) int {
	if a < b {
		return values[a] + sumr6(a+1, b, values)
	}

	return 0
}

// sum7 adds values.
func sum7(values ...int) int {
	return sumr7(0, len(values), values...)
}

// sumr7 recursively adds values over the range [a,b).
func sumr7(a, b int, values ...int) int {
	if a < b {
		return values[a] + sumr7(a+1, b, values...)
	}

	return 0
}
