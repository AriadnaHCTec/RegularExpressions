//go:generate go run gen.go -o extra_test.go

// Package binary128 implements encoding and decoding of IEEE 754 quadruple
// precision floating-point numbers.
//
// https://en.wikipedia.org/wiki/Quadruple-precision_floating-point_format
package binary128

import (
	"fmt"
	"math"
	"math/big"
)

const (
	// precision specifies the number of bits in the mantissa (including the
	// implicit lead bit).
	precision = 113
	// exponent bias.
	bias = 16383
)

// Positive and negative Not-a-Number, infinity and zero.
var (
	// +NaN
	NaN = Float{a: 0x7FFF800000000000, b: 0}
	// -NaN
	NegNaN = Float{a: 0xFFFF800000000000, b: 0}
	// +Inf
	Inf = Float{a: 0x7FFF000000000000, b: 0}
	// -Inf
	NegInf = Float{a: 0xFFFF000000000000, b: 0}
	// +zero
	Zero = Float{a: 0x0000000000000000, b: 0}
	// -zero
	NegZero = Float{a: 0x8000000000000000, b: 0}
)

// Float is a floating-point number in IEEE 754 quadruple precision format.
type Float struct {
	// Sign, exponent and fraction.
	//
	//    1 bit:    sign
	//    15 bits:  exponent
	//    112 bits: fraction
	a uint64
	b uint64
}

// NewFromBits returns the floating-point number corresponding to the IEEE 754
// quadruple precision binary representation.
func NewFromBits(a, b uint64) Float {
	return Float{a: a, b: b}
}

// NewFromFloat32 returns the nearest quadruple precision floating-point number
// for x and the accuracy of the conversion.
func NewFromFloat32(x float32) (Float, big.Accuracy) {
	f, acc := NewFromFloat64(float64(x))
	if acc == big.Exact {