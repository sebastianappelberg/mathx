package mathx

import (
	"math"
	"testing"
)

func TestDivCeil(t *testing.T) {
	test := binaryFuncTest(t, "DivCeil", DivCeil[int])

	// Normal cases
	test(10, 3, 4)   // 10 / 3 = 3.33 → ceil(3.33) = 4
	test(9, 3, 3)    // 9 / 3 = 3.00 → ceil(3.00) = 3
	test(-10, 3, -3) // -10 / 3 = -3.33 → ceil(-3.33) = -3
	test(-10, -3, 4) // -10 / -3 = 3.33 → ceil(3.33) = 4

	// Edge cases
	test(0, 1, 0)   // 0 / 1 = 0 → ceil(0) = 0
	test(1, 1, 1)   // 1 / 1 = 1 → ceil(1) = 1
	test(1, 2, 1)   // 1 / 2 = 0.5 → ceil(0.5) = 1
	test(5, 2, 3)   // 5 / 2 = 2.5 → ceil(2.5) = 3
	test(-5, 2, -2) // -5 / 2 = -2.5 → ceil(-2.5) = -2

	// Large numbers
	test(math.MaxInt64, 2, int(math.Ceil(float64(math.MaxInt64)/2)))
	test(math.MinInt64, 2, int(math.Ceil(float64(math.MinInt64)/2)))
}

func BenchmarkDivCeil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DivCeil(2, 1)
	}
}

func TestDivRound(t *testing.T) {
	test := binaryFuncTest(t, "DivRound", DivRound[int])

	// Normal cases
	test(10, 3, 3)   // 10 / 3 = 3.33 → round(3.33) = 3
	test(9, 3, 3)    // 9 / 3 = 3.00 → round(3.00) = 3
	test(-10, 3, -3) // -10 / 3 = -3.33 → round(-3.33) = -3
	test(-10, -3, 3) // -10 / -3 = 3.33 → round(3.33) = 3

	// Edge cases
	test(0, 1, 0)   // 0 / 1 = 0 → round(0) = 0
	test(1, 1, 1)   // 1 / 1 = 1 → round(1) = 1
	test(1, 2, 1)   // 1 / 2 = 0.5 → round(0.5) = 1
	test(5, 2, 3)   // 5 / 2 = 2.5 → round(2.5) = 3
	test(-5, 2, -3) // -5 / 2 = -2.5 → round(-2.5) = -3

	// Large numbers
	test(math.MaxInt64, 2, int(math.Ceil(float64(math.MaxInt64)/2)))
	test(math.MinInt64, 2, int(math.Ceil(float64(math.MinInt64)/2)))
}

func BenchmarkDivRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DivRound(1, 2)
	}
}

func TestOperandCheck(t *testing.T) {
	// Division by zero should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivCeil(1, 0) did not panic on division by zero")
		}
	}()
	operandCheck(1, 0)

	// NaN cases should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivCeil(NaN, 1) did not panic on NaN")
		}
	}()
	operandCheck[int](int(math.NaN()), 1)

	// Infinity cases should panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivCeil(Inf, 1) did not panic on Inf")
		}
	}()
	operandCheck[int](int(math.Inf(1)), 1)
}
