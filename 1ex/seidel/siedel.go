package seidel

import "math"

func Solve(matrixForm [][]float64, target float64) ([]float64, int) {
	x := make([]float64, len(matrixForm))

	for iteration := 1; ; iteration++ {
		xPrev := make([]float64, len(matrixForm))

		copy(xPrev, x)

		for i := 0; i < len(matrixForm); i++ {
			newX := matrixForm[i][len(matrixForm)]
			for j := 0; j < len(matrixForm); j++ {
				newX += matrixForm[i][j] * x[j]
			}

			x[i] = newX
		}

		dist := 0.0
		for i := 0; i < len(matrixForm); i++ {
			d := x[i] - xPrev[i]
			dist += d * d
		}
		dist = math.Sqrt(dist)

		if dist < target {
			return x, iteration
		}
	}
}
