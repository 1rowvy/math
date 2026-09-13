package utils

func ToForm(matrixA [][]float64, vectorB []float64) [][]float64 {
	matrixForm := make([][]float64, len(matrixA))

	for i := 0; i < len(matrixA); i++ {
		matrixForm[i] = make([]float64, len(matrixA)+1)
		diag := matrixA[i][i]

		for j := 0; j < len(matrixA); j++ {
			if j == i {
				matrixForm[i][j] = 0
			} else {
				matrixForm[i][j] = -matrixA[i][j] / diag
			}
		}

		matrixForm[i][len(matrixA)] = vectorB[i] / diag
	}

	return matrixForm
}
