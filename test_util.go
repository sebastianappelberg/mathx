package mathx

import (
	"math"
	"testing"
)

type binaryFunc[N Number] func(x, y N) N

func binaryFuncTest[N Number](t *testing.T, name string, binaryFunc binaryFunc[N]) func(x, y, want N) {
	return func(x, y, want N) {
		got := binaryFunc(x, y)
		if !equals(got, want) {
			t.Errorf("%s(%v, %v) = %v; want %v", name, x, y, got, want)
		}
	}
}

func equals[N Number](x, y N) bool {
	return math.IsNaN(float64(x)) && math.IsNaN(float64(y)) || x == y
}
