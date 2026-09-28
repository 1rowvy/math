package main

import (
	"fmt"
	"math"
)

// 0.6 -0.35 -0.25
// 0.3 0.25 0.43
// 0.21 -0.45 -0.20
//
// 1.83 0.32 1.91

func transpose(a [][]float64) [][]float64 {
	b := createMatrix()

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[i][j] = a[j][i]
		}
	}

	return b
}

func copyVector(a []float64) []float64 {

	b := make([]float64, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	return b
}

func createMatrix() [][]float64 {
	return [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func mul(m1, m2 [][]float64) [][]float64 {
	m := [][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1); j++ {
			var sum float64

			for k := 0; k < len(m1); k++ {
				sum += m1[i][k] * m2[k][j]
			}

			m[i][j] = sum
		}
	}

	return m
}

func dot(vectorA, vectorB []float64) float64 {
	var sum float64

	for i := 0; i < len(vectorA); i++ {
		sum += vectorA[i] * vectorB[i]
	}

	return sum
}

func gramShmidt(a [][]float64) ([][]float64, [][]float64) {
	n := len(a)
	r := make([][]float64, n)
	t := createMatrix()

	for k := 0; k < n; k++ {

		r[k] = copyVector(a[k])

		for i := 0; i < k; i++ {

			t[i][k] = dot(r[i], a[k]) / dot(r[i], r[i])

			for j := range r[k] {
				r[k][j] -= t[i][k] * r[i][j]
			}
		}
	}

	return r, t
}

func det(r [][]float64) float64 {
	var det float64 = 1

	for i := 0; i < len(r); i++ {
		det *= math.Sqrt(dot(r[i], r[i])
	}

	return det
}

func solve(a [][]float64, r [][]float64, b []float64) []float64 {
	n := len(a)
	x := make([]float64, len(a))
	bc := copyVector(b)

	for i := n - 1; i >= 0; i-- {
		x[i] = dot(r[i], bc) / dot(r[i], a[i])

		for j := range bc {
			bc[j] -= x[i] * a[i][j]
		}
	}

	return x
}

func main() {
	matrixA := [][]float64{
		{0.6, -0.35, -0.25},
		{0.3, 0.25, 0.43},
		{0.21, -0.45, -0.20},
	}

	vectorB := []float64{1.83, 0.32, 1.91}

	transposeA := transpose(matrixA)

	r, t := gramShmidt(transposeA)

	fmt.Println("R:")
	printMatrix(transpose(r))

	fmt.Println("T:")
	printMatrix(t)

	fmt.Println("T*R")
	printMatrix(mul(transpose(r), t))

	fmt.Println(dot(r[0], r[1]), dot(r[0], r[2]), dot(r[1], r[2]))

	fmt.Println("решение")
	fmt.Println(solve(transposeA, r, vectorB))

	fmt.Println("определитель")
	fmt.Println(det(r))
}

func printMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
