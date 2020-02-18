package scratch

import "testing"

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
		if rec := makeFloat32(test.bits); test.exp != rec {
			t.Fatalf("\nexpected %f\nreceived %f\n", test.exp, rec)
		}
	}
}
