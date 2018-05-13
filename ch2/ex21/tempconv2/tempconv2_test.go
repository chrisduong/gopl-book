package tempconv2

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {

	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{68, -205.15},
		{375, 101.85},
	}

	eps := 0.0000001 // acceptable floating point error
	for _, test := range tests {
		k := KToC(test.input)
		if math.Abs(float64(k-test.want)) > eps {
			t.Errorf("KToC(%s): got %s, want %s", test.input, k, test.want)
		}
	}
}
