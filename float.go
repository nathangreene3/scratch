package scratch

import (
	"fmt"
	"math"
)

const (
	// max8BitPow2 (2^8) is the largest power of two possible in 8 bits.
	max8BitPow2 = uint8(1 << 7)

	// max16BitPow2 (2^16) is the largest power of two possible in 16 bits.
	max16BitPow2 = uint16(1 << 15)

	// max32BitPow2 (2^32) is the largest power of two possible in 32 bits.
	max32BitPow2 = uint32(1 << 31)

	// max64BitPow2 (2^64) is the largest power of two possible in 64 bits.
	max64BitPow2 = uint64(1 << 63)
)

func toBits(n interface{}) []byte {
	switch t := n.(type) {
	case int32:
		if t < 0 {
			b32 := to32Bits(uint32(-t))
			b32[31] = 1
			return b32
		}

		return to32Bits(uint32(t))
	case int64:
		if t < 0 {
			b64 := to64Bits(uint64(-t))
			b64[63] = 1
			return b64
		}

		return to64Bits(uint64(t))
	case uint32:
		return to32Bits(t)
	case uint64:
		return to64Bits(t)
	case float32:
		return to32Bits(math.Float32bits(t))
	case float64:
		return to64Bits(math.Float64bits(t))
	default:
		panic(fmt.Sprintf("type %v not supported", t))
	}
}

func to8Bits(n uint8) []byte {
	b := make([]byte, 0, 8)
	for p := uint8(1); 0 < p && p <= max8BitPow2; p <<= 1 {
		if p&n == p {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}

	return b
}

func to16Bits(n uint8) []byte {
	b := make([]byte, 0, 16)
	for p := uint8(1); 0 < p && p <= max8BitPow2; p <<= 1 {
		if p&n == p {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}

	return b
}

func to32Bits(n uint32) []byte {
	b := make([]byte, 0, 32)
	for p := uint32(1); 0 < p && p <= max32BitPow2; p <<= 1 {
		if p&n == p {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}

	return b
}

func to64Bits(n uint64) []byte {
	b := make([]byte, 0, 64)
	for p := uint64(1); 0 < p && p <= max64BitPow2; p <<= 1 {
		if p&n == p {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}

	return b
}

func toFloat32(b32 []byte) float32 {
	return math.Float32frombits(toUInt32(b32))
}

func toFloat64(b64 []byte) float64 {
	return math.Float64frombits(toUInt64(b64))
}

func toInt(b []byte) int {
	switch b[len(b)-1] {
	case 0:
		return int(toUInt(b))
	case 1:
		return -int(toUInt(b[:len(b)-1]))
	default:
		panic("")
	}
}

func toUInt(b []byte) uint {
	var (
		n uint
		p = uint(1)
	)

	for i := range b {
		switch b[i] {
		case 0:
		case 1:
			n += p
		default:
			panic(fmt.Sprintf("%d at index %d is invalid", b[i], i))
		}

		p <<= 1
	}

	return n
}

func toUInt32(b32 []byte) uint32 {
	if len(b32) != 32 {
		panic("invalid length")
	}

	var (
		x uint32
		p = uint32(1)
	)

	for i := range b32 {
		switch b32[i] {
		case 0:
		case 1:
			x += p
		default:
			panic(fmt.Sprintf("%d at index %d is invalid", b32[i], i))
		}

		p <<= 1
	}

	return x
}

func toUInt64(b64 []byte) uint64 {
	if len(b64) != 64 {
		panic("invalid length")
	}

	var (
		x uint64
		p = uint64(1)
	)

	for i := range b64 {
		switch b64[i] {
		case 0:
		case 1:
			x += p
		default:
			panic(fmt.Sprintf("%d at index %d is invalid", b64[i], i))
		}

		p <<= 1
	}

	return x
}

func float64BitsToFloat32Bits(f64 []byte) []byte {
	if len(f64) != 64 {
		panic("invalid length")
	}

	f32 := make([]byte, 32)
	copy(f32[:23], f64[:23])
	copy(f32[23:31], f64[52:60])
	f32[31] = f64[63]
	return f32
}

func float32BitsToFloat64Bits(f32 []byte) []byte {
	if len(f32) != 32 {
		panic("invalid length")
	}

	f64 := make([]byte, 64)
	copy(f64[:23], f32[:23])
	copy(f64[52:60], f32[23:31])
	f64[63] = f32[31]
	return f64
}
