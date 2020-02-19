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
	fs := "11.99999"

	f64, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		t.Fatal(err)
	}

	f32 := float32(f64)
	t.Errorf(
		"\n%s as ...\n"+
			"        float32: %s\n"+
			"        float64: %s\n"+
			"        32 bits: %v\n"+
			"        64 bits: %v\n"+
			"back to float32: %s\n"+
			"back to float64: %s\n",
		fs,
		strconv.FormatFloat(float64(f32), 'f', -1, 32),
		strconv.FormatFloat(float64(f32), 'f', -1, 64),
		toBits(f32),
		toBits(float64(f32)),
		strconv.FormatFloat(float64(toFloat32(toBits(f32))), 'f', -1, 32),
		strconv.FormatFloat(toFloat64(toBits(float64(f32))), 'f', -1, 64),
	)

	t.Errorf(
		"\n%s as ...\n"+
			"        float32: %s\n"+
			"        float64: %s\n"+
			"        32 bits: %v\n"+
			"        64 bits: %v\n"+
			"back to float32: %s\n"+
			"back to float64: %s\n",
		fs,
		strconv.FormatFloat(f64, 'f', -1, 32),
		strconv.FormatFloat(f64, 'f', -1, 64),
		toBits(float32(f64)),
		toBits(f64),
		strconv.FormatFloat(float64(toFloat32(toBits(float32(f64)))), 'f', -1, 32),
		strconv.FormatFloat(toFloat64(toBits(f64)), 'f', -1, 64),
	)
}
