package main

import (
	"reflect"
	"testing"
)

func TestDot(t *testing.T) {
	A1 := NewMatrix64([][]float64{{1.0, 2.0}})
	B1 := NewMatrix64([][]float64{
		{3.0},
		{4.0},
	})
	C1 := NewMatrix64([][]float64{{11.0}})
	C2 := NewMatrix64([][]float64{
		{3.0, 6.0},
		{4.0, 8.0},
	})
	A3 := NewMatrix64([][]float64{{1.2}})
	B3 := NewMatrix64([][]float64{{2.0}})
	C3 := NewMatrix64([][]float64{{2.4}})
	A4 := NewMatrix64([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	})
	B4 := E64(3)
	A5 := Zero64(100, 10000)
	B5 := E64(10000)

	var tests = []struct {
		A *Matrix64
		B *Matrix64
		C *Matrix64
	}{
		{A1, B1, C1},
		{B1, A1, C2},
		{A3, B3, C3},
		{A4, B4, A4},
		{A5, B5, A5},
	}
	for _, test := range tests {
		if result := test.A.Dot(test.B); !reflect.DeepEqual(result, test.C) {
			t.Errorf("Dgemm(A, B)\nA=%v\nB=%v\nwant%v type(%T)\tgot%v type(%T)\n", test.A, test.B, test.C, test.C, result, result)
		}
	}
}

func BenchmarkDot(b *testing.B) {
	//计算1000行1000列矩阵的点乘
	sz := 1000
	ma := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		ma[i] = make([]float64, sz)
	}
	mb := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		mb[i] = make([]float64, sz)
	}
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			ma[i][j] = float64(i)
			mb[i][j] = float64(j)
		}
	}
	var C *Matrix64
	for i := 0; i < b.N; i++ {
		A := NewMatrix64(ma)
		B := NewMatrix64(mb)
		C = A.Dot(B)
	}
	_ = C
}

func TestT(t *testing.T) {
	A1 := NewMatrix64([][]float64{{1.0}})
	A2 := NewMatrix64([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	})
	a2 := NewMatrix64([][]float64{
		{1.0, 4.0},
		{2.0, 5.0},
		{3.0, 6.0},
	})

	var tests = []struct {
		A *Matrix64
		a *Matrix64
	}{
		{A1, A1},
		{A2, a2},
	}

	for _, test := range tests {
		if result := test.A.T(); !reflect.DeepEqual(result, test.a) {
			t.Errorf("T(A)\nA=%v\nwant%v type(%T)\tgot%v type(%T)\n", test.A, test.a, test.a, result, result)
		}
	}
}

func TestScal(t *testing.T) {
	A1 := NewMatrix64([][]float64{{1.0, 2.0}})
	B1 := NewMatrix64([][]float64{{2.0, 4.0}})
	A2 := NewMatrix64([][]float64{
		{3.0, 6.0},
		{4.0, 8.0},
	})
	B2 := NewMatrix64([][]float64{
		{12.0, 24.0},
		{16.0, 32.0},
	})
	A3 := NewMatrix64([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	})
	B3 := NewMatrix64([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	})

	var tests = []struct {
		alpha float64
		A     *Matrix64
		B     *Matrix64
	}{
		{2.0, A1, B1},
		{4.0, A2, B2},
		{1.0, A3, B3},
	}

	for _, test := range tests {
		if result := test.A.Scal(test.alpha); !reflect.DeepEqual(result, test.B) {
			t.Errorf("%v\nScale(%f) want %v but got %v\n", test.A, test.alpha, test.B, result)
		}
	}
}

/*
func BenchmarkMyDot(b *testing.B) {
	//计算1000行1000列矩阵的点乘   三重循环
	sz := 1000
	ma := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		ma[i] = make([]float64, sz)
	}
	mb := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		mb[i] = make([]float64, sz)
	}
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			ma[i][j] = float64(i)
			mb[i][j] = float64(j)
		}
	}

	var mc [][]float64
	for i := 0; i < b.N; i++ {
		mc = myDot(ma, mb)
	}
	_ = mc
}

// 三重循环求矩阵点乘
func myDot(A, B [][]float64) [][]float64 {
	var wg sync.WaitGroup
	sz := len(A)
	mc := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		mc[i] = make([]float64, sz)
	}
	for i := 0; i < sz; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < sz; j++ {
				for k := 0; k < sz; k++ {
					mc[i][j] += A[i][k] * B[k][j]
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return mc
}
*/
