package main

import (
	"reflect"
	"testing"
)

func TestE64(t *testing.T) {
	a1 := NewMatrix64([][]float64{{1.0}})
	a2 := NewMatrix64([][]float64{
		{1.0, 0.0},
		{0.0, 1.0},
	})

	var tests = []struct {
		input  int
		output *Matrix64
	}{
		{1, a1},
		{2, a2},
	}

	for _, test := range tests {
		if result := E64(test.input); !reflect.DeepEqual(result, test.output) {
			t.Errorf("E64(n) error\nn=%d\twant: %v\tgot: %v\n", test.input, test.output, result)
		}
	}
}
