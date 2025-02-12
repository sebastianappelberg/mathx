package mathx

import (
	"math"
	"testing"
)

// TestMin for all numeric types with edge cases
func TestMin(t *testing.T) {
	binaryFuncTest[int](t, "Min", Min)(0, 0, 0)                                       // Zero values
	binaryFuncTest[int](t, "Min", Min)(math.MaxInt, math.MinInt, math.MinInt)         // Max and Min int
	binaryFuncTest[int8](t, "Min", Min)(-128, 127, -128)                              // Min and Max int8
	binaryFuncTest[int16](t, "Min", Min)(-32768, 32767, -32768)                       // Min and Max int16
	binaryFuncTest[int32](t, "Min", Min)(-2147483648, 2147483647, -2147483648)        // Min and Max int32
	binaryFuncTest[int64](t, "Min", Min)(math.MinInt64, math.MaxInt64, math.MinInt64) // Min and Max int64

	binaryFuncTest[uint](t, "Min", Min)(0, math.MaxUint, 0) // Unsigned min/max
	binaryFuncTest[uint8](t, "Min", Min)(255, 0, 0)
	binaryFuncTest[uint16](t, "Min", Min)(65535, 0, 0)
	binaryFuncTest[uint32](t, "Min", Min)(4294967295, 1, 1)
	binaryFuncTest[uint64](t, "Min", Min)(math.MaxUint64, 100, 100)

	binaryFuncTest[float32](t, "Min", Min)(-math.MaxFloat32, math.MaxFloat32, -math.MaxFloat32)
	binaryFuncTest[float64](t, "Min", Min)(-math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64)
	binaryFuncTest[float64](t, "Min", Min)(math.SmallestNonzeroFloat64, 0, 0) // Smallest nonzero value
	binaryFuncTest[float64](t, "Min", Min)(-0.0, 0.0, -0.0)                   // Negative zero case
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Min(1, 2)
	}
}

// TestMax for all numeric types with edge cases
func TestMax(t *testing.T) {
	binaryFuncTest[int](t, "Max", Max)(0, 0, 0)
	binaryFuncTest[int](t, "Max", Max)(math.MaxInt, math.MinInt, math.MaxInt)
	binaryFuncTest[int8](t, "Max", Max)(-128, 127, 127)
	binaryFuncTest[int16](t, "Max", Max)(-32768, 32767, 32767)
	binaryFuncTest[int32](t, "Max", Max)(-2147483648, 2147483647, 2147483647)
	binaryFuncTest[int64](t, "Max", Max)(math.MinInt64, math.MaxInt64, math.MaxInt64)

	binaryFuncTest[uint](t, "Max", Max)(0, math.MaxUint, math.MaxUint)
	binaryFuncTest[uint8](t, "Max", Max)(255, 0, 255)
	binaryFuncTest[uint16](t, "Max", Max)(65535, 0, 65535)
	binaryFuncTest[uint32](t, "Max", Max)(4294967295, 1, 4294967295)
	binaryFuncTest[uint64](t, "Max", Max)(math.MaxUint64, 100, math.MaxUint64)

	binaryFuncTest[float32](t, "Max", Max)(-math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	binaryFuncTest[float64](t, "Max", Max)(-math.MaxFloat64, math.MaxFloat64, math.MaxFloat64)
	binaryFuncTest[float64](t, "Max", Max)(math.SmallestNonzeroFloat64, 0, math.SmallestNonzeroFloat64)
	binaryFuncTest[float64](t, "Max", Max)(-0.0, 0.0, 0.0) // Negative zero case
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max(1, 2)
	}
}

func TestDim(t *testing.T) {
	binaryFuncTest[int](t, "Dim", Dim)(5, 3, 2)
	binaryFuncTest[int](t, "Dim", Dim)(10, 10, 0) // Identical values
	binaryFuncTest[int](t, "Dim", Dim)(-5, -3, 0) // Negative values
	binaryFuncTest[int](t, "Dim", Dim)(100, 50, 50)

	binaryFuncTest[uint](t, "Dim", Dim)(10, 10, 0)
	binaryFuncTest[uint8](t, "Dim", Dim)(255, 0, 255)
	binaryFuncTest[uint16](t, "Dim", Dim)(65535, 0, 65535)
	binaryFuncTest[uint32](t, "Dim", Dim)(4294967295, 1, 4294967294)
	binaryFuncTest[uint64](t, "Dim", Dim)(math.MaxUint64, 100, math.MaxUint64-100)

	binaryFuncTest[float32](t, "Dim", Dim)(5.5, 2.2, 3.3)
	binaryFuncTest[float64](t, "Dim", Dim)(-10.1, -20.2, 10.1)
	binaryFuncTest[float64](t, "Dim", Dim)(math.MaxFloat64, 0, math.MaxFloat64)    // Large float case
	binaryFuncTest[float64](t, "Dim", Dim)(0, math.SmallestNonzeroFloat64, 0)      // Smallest float case
	binaryFuncTest[float64](t, "Dim", Dim)(-0.0, 0.0, 0.0)                         // Negative zero case
	binaryFuncTest[float64](t, "Dim", Dim)(math.Inf(1), math.Inf(1), math.NaN())   // +Inf, +Inf = NaN edge case.
	binaryFuncTest[float64](t, "Dim", Dim)(math.Inf(-1), math.Inf(-1), math.NaN()) // -Inf, -Inf = NaN edge case.
	binaryFuncTest[float64](t, "Dim", Dim)(1, math.NaN(), math.NaN())              // x, NaN = NaN edge case.
}

func TestDimIntOverflow(t *testing.T) {
	testInt8 := binaryFuncTest[int8](t, "Dim", Dim)
	testInt8(math.MaxInt8, -1, 0)
	testInt8(math.MaxInt8, math.MinInt8, 0)
	testInt8(math.MinInt8, 1, math.MaxInt8)
	testInt8(math.MinInt8, math.MaxInt8, 1)

	testInt16 := binaryFuncTest[int16](t, "Dim", Dim)
	testInt16(math.MaxInt16, -1, 0)
	testInt16(math.MaxInt16, math.MinInt16, 0)
	testInt16(math.MinInt16, 1, math.MaxInt16)
	testInt16(math.MinInt16, math.MaxInt16, 1)

	testInt32 := binaryFuncTest[int32](t, "Dim", Dim)
	testInt32(math.MaxInt32, -1, 0)
	testInt32(math.MaxInt32, math.MinInt32, 0)
	testInt32(math.MinInt32, 1, math.MaxInt32)
	testInt32(math.MinInt32, math.MaxInt32, 1)

	testInt := binaryFuncTest[int](t, "Dim", Dim)
	testInt(math.MaxInt, -1, 0)
	testInt(math.MaxInt, math.MinInt, 0)
	testInt(math.MinInt, 1, math.MaxInt)
	testInt(math.MinInt, math.MaxInt, 1)

	testInt64 := binaryFuncTest[int64](t, "Dim", Dim)
	testInt64(math.MaxInt64, -1, 0)
	testInt64(math.MaxInt64, math.MinInt64, 0)
	testInt64(math.MinInt64, 1, math.MaxInt64)
	testInt64(math.MinInt64, math.MaxInt64, 1)
}

func BenchmarkDim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Dim(5, 1)
	}
}
