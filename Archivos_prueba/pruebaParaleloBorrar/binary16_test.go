package binary16

import (
	"math"
	"math/big"
	"testing"
)

func TestNewFromBits(t *testing.T) {
	golden := []struct {
		bits uint16
		want float64
	}{
		// Special numbers.
		// 0 11111 1000000000 = +NaN
		{bits: 0x7E00, want: math.NaN()},
		// -NaN
		// 1 11111 1000000000 = -NaN
		{bits: 0xFE00, want: -math.NaN()},

		// from: https://en.wikipedia.org/wiki/Half-precision_floating-point_format#Half_precision_examples

		// 0 01111 0000000000 = 1
		{bits: 0x3C00, want: 1},
		// 0 01111 0000000001 = 1 + 2^(-10) = 1.0009765625 (next smallest float after 1)
		{bits: 0x3C01, want: 1.0009765625},
		// 1 10000 0000000000 = -2
		{bits: 0xC000, want: -2},
		// 0 11110 1111111111 = 65504 (max half precision)
		{bits: 0x7BFF, want: 65504},
		// 0 00001 0000000000 = 2^(-14) ~= 6.10352 * 10^(-5) (minimum positive normal)
		{bits: 0x0400, want: math.Pow(2, -14)},
		// 0 00000 0000000001 = 2^(-24) ~= 5.96046 * 10^(-8) (minimum positive subnormal)
		{bits: 0x0001, want: math.Pow(2, -24)},
		// 0 00000 0000000000 = 0
		{bits: 0x0000, want: 0},
		// 1 00000 0000000000 = −0
		{bits: 0x8000, want: math.Copysign(0, -1)},
		// 0 11111 0000000000 = infinity
		{bits: 0x7C00, want: math.Inf(1)},
		// 1 11111 0000000000 = -infinity
		{bits: 0xFC00, want: math.Inf(-1)},
		// 0 01101 0101010101 = 0.333251953125 ~= 1/3
		{bits: 0x3555, want: 0.333251953125},

		// from: https://reviews.llvm.org/rL237161

		// Normalized numbers.
		// 0 01110 0000000000 = 0.5
		{bits: 0x3800, want: 0.5},
		// 1 01110 0000000000 = -0.5
		{bits: 0xB800, want: -0.5},
		// 0 01111 1000000000 = 1.5
		{bits: 0x3E00, want: 1.5},
		// 1 01111 1000000000 = -1.5
		{bits: 0xBE00, want: -1.5},
		// 0 10000 0100000000 = 2.5
		{bits: 0x4100, want: 2.5},
		// 1 10000 0100000000 = -2.5