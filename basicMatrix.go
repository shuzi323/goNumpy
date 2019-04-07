package main

/* float64 */

// Zero64 返回m行n列的矩阵  float64
func Zero64(m, n int) *Matrix64 {
	mat := make([]float64, m*n)
	return &Matrix64{mat: mat, row: m, col: n}
}

//One64 返回m行n列全是1.0的矩阵	float64
func One64(m, n int) *Matrix64 {
	A := Zero64(m, n)
	for i := 0; i < m*n; i++ {
		A.mat[i] = 1.0
	}
	return A
}

// E64 返回m行的单位矩阵
func E64(m int) *Matrix64 {
	A := Zero64(m, m)
	for i := 0; i < m; i++ {
		A.mat[i+i*m] = 1.0
	}
	return A
}

/* float32 */

// Zero32 返回m行n列的矩阵  float32
func Zero32(m, n int) *Matrix32 {
	mat := make([]float32, m*n)
	return &Matrix32{mat: mat, row: m, col: n}
}

//One32 返回m行n列全是1.0的矩阵	float32
func One32(m, n int) *Matrix32 {
	A := Zero32(m, n)
	for i := 0; i < m*n; i++ {
		A.mat[i] = 1.0
	}
	return A
}

// E32 返回m行的单位矩阵
func E32(m int) *Matrix32 {
	A := Zero32(m, m)
	for i := 0; i < m; i++ {
		A.mat[i+i*m] = 1.0
	}
	return A
}
