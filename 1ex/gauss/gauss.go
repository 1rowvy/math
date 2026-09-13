package gauss

func Solve(matrixA [][]float64, vectorB []float64) []float64 {

	a := make([][]float64, len(matrixA))
	b := make([]float64, len(matrixA))

	for i := 0; i < len(matrixA); i++ {
		a[i] = make([]float64, len(matrixA))
		copy(a[i], matrixA[i])

		b[i] = vectorB[i]
	}

	for k := 0; k < len(matrixA); k++ {
		for i := k + 1; i < len(matrixA); i++ {
			factor := a[i][k] / a[k][k]
			for j := k; j < len(matrixA); j++ {
				a[i][j] -= factor * a[k][j]
			}
			b[i] -= factor * b[k]
		}
	}

	x := make([]float64, len(matrixA))
	for i := len(matrixA) - 1; i >= 0; i-- {
		sum := b[i]
		for j := i + 1; j < len(matrixA); j++ {
			sum -= a[i][j] * x[j]
		}
		x[i] = sum / a[i][i]
	}

	return x

}
