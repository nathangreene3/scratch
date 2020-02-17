package scratch

import "testing"

// newTestInts returns [1, 2, ..., n]
func newTestInts(n int) []int {
	ints := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		ints = append(ints, i)
	}
	return ints
}

func TestSums(t *testing.T) {
	for n := 1; n <= 5; n++ {
		var (
			values = newTestInts(n)
			exp    = (n*n + n) / 2
			rec    = map[string]int{
				"sum0": sum0(values),
				"sum1": sum1(values),
				"sum2": sum2(values...),
				"sum3": sum3(values...),
				"sum4": sum4(values),
				"sum5": sum5(values...),
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
	values := newTestInts(256)
	benchmarkSum(b, values)
	benchmarkSum0(b, values)
	benchmarkSum1(b, values)
	benchmarkSum2(b, values)
	benchmarkSum3(b, values)
	benchmarkSum4(b, values)
	benchmarkSum5(b, values)
}

func benchmarkSum(b *testing.B, values []int) {
	b.Run(
		"sum0",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum(values)
			}
		},
	)
}

func benchmarkSum0(b *testing.B, values []int) {
	b.Run(
		"sum0",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum0(values)
			}
		},
	)
}

func benchmarkSum1(b *testing.B, values []int) {
	b.Run(
		"sum1",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum1(values)
			}
		},
	)
}

func benchmarkSum2(b *testing.B, values []int) {
	b.Run(
		"sum2",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum2(values...)
			}
		},
	)
}

func benchmarkSum3(b *testing.B, values []int) {
	b.Run(
		"sum3",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum3(values...)
			}
		},
	)
}

func benchmarkSum4(b *testing.B, values []int) {
	b.Run(
		"sum4",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum4(values)
			}
		},
	)
}

func benchmarkSum5(b *testing.B, values []int) {
	b.Run(
		"sum5",
		func(b1 *testing.B) {
			for j := 0; j < b1.N; j++ {
				_ = sum5(values...)
			}
		},
	)
}
