package scratch

import (
	"strconv"
	"testing"
)

// intsModN returns [0, 1, ..., n-1]
func intsModN(n int) []int {
	ints := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}

	return ints
}

func TestSums(t *testing.T) {
	for n := 0; n < 8; n++ {
		var (
			values = intsModN(n)
			exp    = (n*n - n) / 2 // 0 + 1 + ... + (n-1) = (n^2 - n)/2
			rec    = map[string]int{
				"sum0": sum0(values),
				"sum1": sum1(values),
				"sum2": sum2(values),
				"sum3": sum3(values...),
				"sum4": sum4(values...),
				"sum5": sum5(values...),
				"sum6": sum6(values),
				"sum7": sum7(values...),
			}
		)

		for f, r := range rec {
			if exp != r {
				t.Fatalf("\n   given %d\nexpected %d from function '%s'\nreceived %d\n", n, exp, f, r)
			}
		}
	}
}

func BenchmarkSums(b *testing.B) {
	// Linear scale
	maxLinear := 8
	for n := 0; n < maxLinear; n++ {
		benchmarkSumSlice(b, "sum0", sum0, intsModN(n))
	}

	for n := 0; n < maxLinear; n++ {
		benchmarkSumSlice(b, "sum1", sum1, intsModN(n))
	}

	for n := 0; n < maxLinear; n++ {
		benchmarkSumSlice(b, "sum2", sum2, intsModN(n))
	}

	for n := 0; n < maxLinear; n++ {
		benchmarkSumVariadic(b, "sum3", sum3, intsModN(n))
	}

	for n := 0; n < maxLinear; n++ {
		benchmarkSumVariadic(b, "sum4", sum4, intsModN(n))
	}

	for n := 0; n < maxLinear; n++ {
		benchmarkSumVariadic(b, "sum5", sum5, intsModN(n))
	}

	for n := 0; n < 8; n++ {
		benchmarkSumSlice(b, "sum6", sum6, intsModN(n))
	}

	for n := 0; n < 8; n++ {
		benchmarkSumVariadic(b, "sum7", sum7, intsModN(n))
	}

	// Exponential scale
	maxExponential := 256
	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumSlice(b, "sum0", sum0, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumSlice(b, "sum1", sum1, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumSlice(b, "sum2", sum2, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumVariadic(b, "sum3", sum3, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumVariadic(b, "sum4", sum4, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumVariadic(b, "sum5", sum5, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumSlice(b, "sum6", sum6, intsModN(n))
	}

	for n := 1; n <= maxExponential; n <<= 1 {
		benchmarkSumVariadic(b, "sum7", sum7, intsModN(n))
	}
}

func benchmarkSumSlice(b *testing.B, fnName string, sum func(values []int) int, values []int) {
	b.Run(
		fnName+" len "+strconv.Itoa(len(values)),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sum(values)
			}
		},
	)
}

func benchmarkSumVariadic(b *testing.B, fnName string, sum func(values ...int) int, values []int) {
	b.Run(
		fnName+" len "+strconv.Itoa(len(values)),
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sum(values...)
			}
		},
	)
}
