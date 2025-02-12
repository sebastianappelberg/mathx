package mathx

import "math"

// DivCeil returns the least integer value greater than or equal to the quotient of x and y.
//
// The function will panic in these cases:
//
//	x = 0
//	x or y = NaN
//	x or y = ±Inf
func DivCeil[N Number](x, y N) int {
	operandCheck(x, y)
	return int(math.Ceil(float64(x) / float64(y)))
}

// DivRound returns the nearest integer to the quotient of x and y, rounding half away from zero.
//
// The function will panic in these cases:
//
//	x = 0
//	x or y = NaN
//	x or y = ±Inf
func DivRound[N Number](x, y N) int {
	operandCheck(x, y)
	return int(math.Round(float64(x) / float64(y)))
}

func operandCheck[N Number](x, y N) {
	if y == 0 {
		panic("division by zero")
	}
	if math.IsNaN(float64(x)) {
		panic("operand cannot be NaN")
	}
	if math.IsInf(float64(x), 1) || math.IsInf(float64(x), -1) {
		panic("operand cannot be Inf")
	}
}
