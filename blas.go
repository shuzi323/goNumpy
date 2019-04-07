package main

//#cgo CFLAGS: -I./OpenBLAS-Win64-int64/include/
//#cgo LDFLAGS: ${SRCDIR}/OpenBLAS-Win64-int64/lib/libopenblas.a -lpthread
//
//#include "cblas.h"
import "C"
import (
	"unsafe"
)

const (
	//Trans 不需要转置
	Trans = C.CblasTrans
	//NoTrans 需要转置
	NoTrans = C.CblasNoTrans
)

//Dgemm 用于float64的矩阵乘法 α*A·B + βD, 无返回值，结果在D中
func Dgemm(transA, transB uint32, alpha float64, A, B *Matrix64, beta float64, D *Matrix64) {
	m := A.row
	n := B.col
	k := A.col
	if transA == Trans {
		m = A.col
		k = A.row
	}
	if transB == Trans {
		n = B.row
	}
	if A.row != D.row || B.col != D.col {
		panic("matrix A, B and D not match!")
	}
	if A.col != B.row {
		panic("matrix Col A != Row B!")
	}
	C.cblas_dgemm(C.CblasRowMajor, transA, transB, C.longlong(m), C.longlong(n), C.longlong(k), C.double(alpha), (*C.double)(unsafe.Pointer(&A.mat[0])), C.longlong(A.col), (*C.double)(unsafe.Pointer(&B.mat[0])), C.longlong(B.col), C.double(beta), (*C.double)(unsafe.Pointer(&D.mat[0])), C.longlong(D.col))
}

//Sgemm 用于float32的矩阵乘法 α*A·B + βD, 无返回值，结果在D中
func Sgemm(transA, transB uint32, alpha float32, A, B *Matrix32, beta float32, D *Matrix32) {
	m := A.row
	n := B.col
	k := A.col
	if transA == Trans {
		m = A.col
		k = A.row
	}
	if transB == Trans {
		n = B.row
	}
	if A.row != D.row || B.col != D.col {
		panic("matrix A, B and D not match!")
	}
	if A.col != B.row {
		panic("matrix Col A != Row B!")
	}
	C.cblas_sgemm(C.CblasRowMajor, transA, transB, C.longlong(m), C.longlong(n), C.longlong(k), C.float(alpha), (*C.float)(unsafe.Pointer(&A.mat[0])), C.longlong(A.col), (*C.float)(unsafe.Pointer(&B.mat[0])), C.longlong(B.col), C.float(beta), (*C.float)(unsafe.Pointer(&D.mat[0])), C.longlong(D.col))
}

func dgemm(transA, transB uint32, m, n, k int, alpha float64, A *Matrix64, lda int, B *Matrix64, ldb int, beta float64, D *Matrix64, ldc int) {
	C.cblas_dgemm(C.CblasRowMajor, transA, transB, C.longlong(m), C.longlong(n), C.longlong(k), C.double(alpha), (*C.double)(unsafe.Pointer(&A.mat[0])), C.longlong(lda), (*C.double)(unsafe.Pointer(&B.mat[0])), C.longlong(ldb), C.double(beta), (*C.double)(unsafe.Pointer(&D.mat[0])), C.longlong(ldc))
}

func sgemm(transA, transB uint32, m, n, k int, alpha float32, A *Matrix32, lda int, B *Matrix32, ldb int, beta float32, D *Matrix32, ldc int) {
	C.cblas_sgemm(C.CblasRowMajor, transA, transB, C.longlong(m), C.longlong(n), C.longlong(k), C.float(alpha), (*C.float)(unsafe.Pointer(&A.mat[0])), C.longlong(lda), (*C.float)(unsafe.Pointer(&B.mat[0])), C.longlong(ldb), C.float(beta), (*C.float)(unsafe.Pointer(&D.mat[0])), C.longlong(ldc))
}
