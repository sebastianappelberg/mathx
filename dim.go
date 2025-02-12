package mathx

// Min uses the min built-in function which returns the smallest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If N is a floating-point type and any of the arguments are NaNs,
// min will return NaN.
func Min[N Number](x, y N) N {
	return min(x, y)
}

// Max uses the max built-in function which returns the largest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If N is a floating-point type and any of the arguments are NaNs,
// max will return NaN.
func Max[N Number](x, y N) N {
	return max(x, y)
}

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//
//	Dim(MaxInteger, -x) = 0.
//	Dim(MinInteger, x) = MinInteger + 1 - x.
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim[N Number](x, y N) N {
	v := x - y
	if v <= 0 {
		return 0
	}
	// v is positive or NaN
	return v
}
