package simpleIteration

import "math"

func Solve(matrixForm [][]float64, target float64) ([]float64, int) {

	x := make([]float64, len(matrixForm))

	for iteration := 1; ; iteration++ {
		xNext := make([]float64, len(matrixForm))

		for i := 0; i < len(matrixForm); i++ {
			newX := matrixForm[i][len(matrixForm)]
			for j := 0; j < len(matrixForm); j++ {
				newX += matrixForm[i][j] * x[j]
			}

			xNext[i] = newX
		}

		dist := 0.0
		for i := 0; i < len(matrixForm); i++ {
			d := xNext[i] - x[i]
			dist += d * d
		}
		dist = math.Sqrt(dist)

		if dist < target {
			return xNext, iteration
		}

		x = xNext
	}
}

func metric1(matrixForm [][]float64) float64 {
	max := 0.0

	for i := 0; i < len(matrixForm); i++ {
		sum := 0.0
		for j := 0; j < len(matrixForm); j++ {
			sum += math.Abs(matrixForm[i][j])
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

func metric2(matrixForm [][]float64) float64 {
	max := 0.0

	for j := 0; j < len(matrixForm); j++ {
		sum := 0.0
		for i := 0; i < len(matrixForm); i++ {
			sum += math.Abs(matrixForm[i][j])
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

func metric3(matrixForm [][]float64) float64 {
	sum := 0.0

	for i := 0; i < len(matrixForm); i++ {
		for j := 0; j < len(matrixForm); j++ {
			sum += matrixForm[i][j] * matrixForm[i][j]
		}
	}

	return math.Sqrt(sum)
}

func CheckConvergence(matrixForm [][]float64) (bool, float64, float64, float64, float64) {
	m1 := metric1(matrixForm)
	m2 := metric2(matrixForm)
	m3 := metric3(matrixForm)

	convergent := m1 < 1 || m2 < 1 || m3 < 1

	return convergent, m1, m2, m3, min(m1, m2, m3)
}

func CalcTarget(alpha, eps float64) float64 {
	target := eps * (1 - alpha) / alpha

	return target
}
