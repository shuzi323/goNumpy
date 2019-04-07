package main

import (
	"log"
)

//Matrix64 float64 的矩阵
type Matrix64 struct {
	mat []float64 //将矩阵存为slice
	row int       //行
	col int       //列
}

//NewMatrix64 返回内置的矩阵格式， float64
func NewMatrix64(A [][]float64) *Matrix64 {
	m, n, mat := matrixToArray64(A)
	return &Matrix64{mat: mat, row: m, col: n}
}

func matrixToArray64(A [][]float64) (m, n int, mat []float64) {
	m = len(A)
	n = len(A[0])
	a := make([]float64, m*n)
	x := 0
	for i := 0; i < m; i++ {
		if len(A[i]) != n {
			log.Println(A)
			panic("matrix err!")
		}
		for j := 0; j < n; j++ {
			a[x] = A[i][j]
			x++
		}
	}
	return m, n, a
}

//Dot 计算矩阵的点乘
func (A *Matrix64) Dot(B *Matrix64) *Matrix64 {
	d := make([]float64, A.row*B.col)
	D := &Matrix64{mat: d, row: A.row, col: B.col}
	Dgemm(NoTrans, NoTrans, 1.0, A, B, 0.0, D)
	return D
}

//T 返回矩阵的转置
func (A *Matrix64) T() *Matrix64 {
	d := make([]float64, A.row*A.col)
	for i := 0; i < A.row; i++ {
		for j := 0; j < A.col; j++ {
			d[j*A.row+i] = A.mat[i*A.col+j]
		}
	}
	return &Matrix64{mat: d, row: A.col, col: A.row}
}

// Scal 返回α*A
func (A *Matrix64) Scal(alpha float64) *Matrix64 {
	d := make([]float64, A.row*A.col)
	D := &Matrix64{mat: d, row: A.row, col: A.col}
	for i := 0; i < len(d); i++ {
		d[i] = alpha * A.mat[i]
	}
	return D
}

// Reshape A.row, A.col = m, n
func (A *Matrix64) Reshape(m, n int) {
	if n < 0 || m < 0 || (m+n) != (A.row+A.col) {
		panic("Reshape error; invalid m, n!")
	}
	A.row, A.col = m, n
}

// Get 获取row行col列的值。从0行0列开始
func (A *Matrix64) Get(row, col int) float64 {
	return A.mat[row*A.row+col]
}

// Trace 方阵的迹
func (A *Matrix64) Trace() float64 {
	if A.row != A.col {
		log.Println(A)
		panic("not square!")
	}
	var tr float64
	for i := 0; i < A.row; i++ {
		tr += A.Get(i, i)
	}
	return tr
}

//Matrix32 float32 的矩阵
type Matrix32 struct {
	mat []float32 //将矩阵存为slice
	row int       //行
	col int       //列
}

//NewMatrix32 返回内置的矩阵格式， float32
func NewMatrix32(A [][]float32) *Matrix32 {
	m, n, mat := matrixToArray32(A)
	return &Matrix32{mat: mat, row: m, col: n}
}

func matrixToArray32(A [][]float32) (m, n int, mat []float32) {
	m = len(A)
	n = len(A[0])
	a := make([]float32, m*n)
	x := 0
	for i := 0; i < m; i++ {
		if len(A[i]) != n {
			log.Println(A)
			panic("matrix err!")
		}
		for j := 0; j < n; j++ {
			a[x] = A[i][j]
			x++
		}
	}
	return m, n, a
}

//Dot 计算矩阵的点乘
func (A *Matrix32) Dot(B *Matrix32) *Matrix32 {
	d := make([]float32, A.row*B.col)
	D := &Matrix32{mat: d, row: A.row, col: B.col}
	Sgemm(NoTrans, NoTrans, 1.0, A, B, 0.0, D)
	return D
}

//T 返回矩阵的转置
func (A *Matrix32) T() *Matrix32 {
	d := make([]float32, A.row*A.col)
	for i := 0; i < A.row; i++ {
		for j := 0; j < A.col; j++ {
			d[j*A.row+i] = A.mat[i*A.col+j]
		}
	}
	return &Matrix32{mat: d, row: A.col, col: A.row}
}

// Get 获取row行col列的值。从0行0列开始
func (A *Matrix32) Get(row, col int) float32 {
	return A.mat[row*A.row+col]
}

// Trace 方阵的迹
func (A *Matrix32) Trace() float32 {
	if A.row != A.col {
		log.Println(A)
		panic("not square!")
	}
	var tr float32
	for i := 0; i < A.row; i++ {
		tr += A.Get(i, i)
	}
	return tr
}
