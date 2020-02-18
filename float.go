package scratch

import "math"

func makeFloat32(bits []byte) float32 {
	f := float32(1)
	for _, b := range bits[:23] {
		f += float32(b) * float32(math.Pow(2, -float64(b)-1))
	}

	f *= float32(math.Pow(2, float64(toInt(bits[23:31])-127)))
	if bits[31] == 0 {
		return f
	}

	return -f
}

func toInt(bits []byte) int {
	var (
		x int
		p = 1
	)

	for _, b := range bits {
		if b == 1 {
			x += p
		}

		p <<= 1
	}

	return x
}

func makeFloat(bits []byte) float64 {
	f := float64(1)
	for _, b := range bits[:52] {
		f += float64(b) * math.Pow(2, -float64(b)-1)
	}

	f *= math.Pow(2, float64(toInt(bits[52:63]))-1023)
	if bits[63] == 0 {
		return f
	}

	return -f
}

func convert32To64(bits []byte) []byte {
	bigBits := make([]byte, 64)

	return bigBits
}
