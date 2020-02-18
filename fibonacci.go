package scratch

func fibonacci(n int) int {
	a0, a1 := 1, 1
	for ; 1 < n; n-- {
		a0, a1 = a1, a0+a1
	}

	return a1
}

func fibonacci1(n int) int {
	return 0
}

// 11.99999 --> 11.999990463256836
