package scratch

import "testing"

import "strconv"

func TestMakeFloat32(t *testing.T) {
	tests := []struct {
		bits []byte
		exp  float32
	}{
		{
			bits: []byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, // Fraction (23 bits)
				0, 0, 1, 1, 1, 1, 1, 0, // Exponent (8 bits)
				0, // Sign (1 bit)
			},
			exp: 0.15625,
		},
	}

	for _, test := range tests {
		if rec := toFloat32(test.bits); test.exp != rec {
			t.Fatalf("\nexpected %f\nreceived %f\n", test.exp, rec)
		}
	}
}

func TestToBits(t *testing.T) {
	var (
		fs       = "0.9"
		f64, err = strconv.ParseFloat(fs, 64)
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Fatalf("\n%s as ...\n"+
		"float32: %s\n"+
		"float64: %s\n"+
		"32 bits: %v\n"+
		"64 bits: %v\n",
		fs,
		strconv.FormatFloat(f64, 'f', -1, 32),
		strconv.FormatFloat(f64, 'f', -1, 64),
		toBits(float32(f64)),
		toBits(f64),
	)
}

func TestInt(t *testing.T) {
	n := int(uint(1<<32 - 1))
	t.Fatal(int(int32(n)) == n, int(int64(n)) == n)
}
