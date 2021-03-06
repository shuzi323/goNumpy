package main

import (
	"fmt"
)

func main() {
	var str string
	str = "hello"
	str = str + " world!"
	fmt.Println(str)
	a := []float32{1.0, 2.3, 5.0, 45.3}
	fmt.Printf("%g\n", a)
	fmt.Println(a)
	B := [][]float64{
		{1.0, 2.0, 1.0, 4.3, -1.9, 4.3, -1.9, 5.0, 45.3},
		{-3.0, 4.0, -1.0, 4.3, -1.9, 4.3, -1.9, 5.0, 45.3},
		{-3.5, 4.3, -1.9, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{1.0, 2.0, 1.0, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{-3.0, 4.0, -1.0, 4.3, -1.9, 4.3, -1.9, 4.0, -1.0},
		{-3.5, 4.3, -1.9, 4.3, -1.9, 4.3, -1.9, 4.0, -1.0},
		{-3.5, 4.3, -1.9, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{1.0, 2.0, 1.0, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{-3.5, 4.3, -1.9, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{1.0, 2.0, 1.0, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{-3.0, 4.0, -1.0, 4.3, -1.9, 4.3, -1.9, 5.0, 45.3},
		{-3.5, 4.3, -1.9, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{1.0, 2.0, 1.0, 4.3, -1.9, 4.3, -1.9, 2.3, 5.0},
		{-243.234, -4.0, -1.0, 4.3, -1.9, 4.3, -1.9, 4.0, -1.0},
	}
	b := NewMatrix64(B)
	fmt.Println("b:")
	fmt.Println(b)
}
