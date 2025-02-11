package mathx

import (
	"math"
	"testing"
)

// Generic helper functions for testing
func runMinTest[T Number](t *testing.T, x, y, expected T) {
	result := Min(x, y)
	if result != expected {
		t.Errorf("Min(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

// TestMin for all numeric types with edge cases
func TestMin(t *testing.T) {
	runMinTest[int](t, 0, 0, 0)                                       // Zero values
	runMinTest[int](t, math.MaxInt, math.MinInt, math.MinInt)         // Max and Min int
	runMinTest[int8](t, -128, 127, -128)                              // Min and Max int8
	runMinTest[int16](t, -32768, 32767, -32768)                       // Min and Max int16
	runMinTest[int32](t, -2147483648, 2147483647, -2147483648)        // Min and Max int32
	runMinTest[int64](t, math.MinInt64, math.MaxInt64, math.MinInt64) // Min and Max int64

	runMinTest[uint](t, 0, math.MaxUint, 0) // Unsigned min/max
	runMinTest[uint8](t, 255, 0, 0)
	runMinTest[uint16](t, 65535, 0, 0)
	runMinTest[uint32](t, 4294967295, 1, 1)
	runMinTest[uint64](t, math.MaxUint64, 100, 100)

	runMinTest[float32](t, -math.MaxFloat32, math.MaxFloat32, -math.MaxFloat32)
	runMinTest[float64](t, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64)
	runMinTest[float64](t, math.SmallestNonzeroFloat64, 0, 0) // Smallest nonzero value
	runMinTest[float64](t, -0.0, 0.0, -0.0)                   // Negative zero case
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Min(1, 2)
	}
}

func runMaxTest[T Number](t *testing.T, x, y, expected T) {
	result := Max(x, y)
	if result != expected {
		t.Errorf("Max(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

// TestMax for all numeric types with edge cases
func TestMax(t *testing.T) {
	runMaxTest[int](t, 0, 0, 0)
	runMaxTest[int](t, math.MaxInt, math.MinInt, math.MaxInt)
	runMaxTest[int8](t, -128, 127, 127)
	runMaxTest[int16](t, -32768, 32767, 32767)
	runMaxTest[int32](t, -2147483648, 2147483647, 2147483647)
	runMaxTest[int64](t, math.MinInt64, math.MaxInt64, math.MaxInt64)

	runMaxTest[uint](t, 0, math.MaxUint, math.MaxUint)
	runMaxTest[uint8](t, 255, 0, 255)
	runMaxTest[uint16](t, 65535, 0, 65535)
	runMaxTest[uint32](t, 4294967295, 1, 4294967295)
	runMaxTest[uint64](t, math.MaxUint64, 100, math.MaxUint64)

	runMaxTest[float32](t, -math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	runMaxTest[float64](t, -math.MaxFloat64, math.MaxFloat64, math.MaxFloat64)
	runMaxTest[float64](t, math.SmallestNonzeroFloat64, 0, math.SmallestNonzeroFloat64)
	runMaxTest[float64](t, -0.0, 0.0, 0.0) // Negative zero case
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max(1, 2)
	}
}

func runDimTest[T Number](t *testing.T, x, y, expected T) {
	result := Dim(x, y)
	if math.IsNaN(float64(expected)) && !math.IsNaN(float64(result)) || (!math.IsNaN(float64(expected)) && result != expected) {
		t.Errorf("Dim(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

// TestDim for all numeric types with edge cases
func TestDim(t *testing.T) {
	runDimTest[int](t, 5, 3, 2)
	runDimTest[int](t, 10, 10, 0) // Identical values
	runDimTest[int](t, -5, -3, 0) // Negative values
	runDimTest[int](t, 100, 50, 50)
	runDimTest[int64](t, math.MaxInt64, math.MinInt64, 0) // Int64 overflow edge case.

	runDimTest[uint](t, 10, 10, 0)
	runDimTest[uint8](t, 255, 0, 255)
	runDimTest[uint16](t, 65535, 0, 65535)
	runDimTest[uint32](t, 4294967295, 1, 4294967294)
	runDimTest[uint64](t, math.MaxUint64, 100, math.MaxUint64-100)

	runDimTest[float32](t, 5.5, 2.2, 3.3)
	runDimTest[float64](t, -10.1, -20.2, 10.1)
	runDimTest[float64](t, math.MaxFloat64, 0, math.MaxFloat64)    // Large float case
	runDimTest[float64](t, 0, math.SmallestNonzeroFloat64, 0)      // Smallest float case
	runDimTest[float64](t, -0.0, 0.0, 0.0)                         // Negative zero case
	runDimTest[float64](t, math.Inf(1), math.Inf(1), math.NaN())   // +Inf, +Inf = NaN edge case.
	runDimTest[float64](t, math.Inf(-1), math.Inf(-1), math.NaN()) // -Inf, -Inf = NaN edge case.
	runDimTest[float64](t, 1, math.NaN(), math.NaN())              // x, NaN = NaN edge case.
}

func BenchmarkDim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Dim(5, 1)
	}
}
