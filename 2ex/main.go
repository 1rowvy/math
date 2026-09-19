package main

import "fmt"

// 1 4 0
// 0 1 0
// 0 3 1

func SolveLU(l, u [][]float64, b []float64) []float64 {
	x := make([]float64, len(b))

	for i := 0; i < len(b); i++ {
		sum := b[i]

		for k := 0; k < i; k++ {
			sum -= l[i][k] * x[k]
		}

		x[i] = sum / l[i][i]
	}

	for i := len(b) - 1; i >= 0; i-- {
		sum := 0.0

		for k := i + 1; k < len(b); k++ {
			sum += u[i][k] * x[k]
		}

		x[i] = (x[i] - sum) / u[i][i]
	}

	return x
}

func SolveLDU(l, d, u1 [][]float64, b []float64) []float64 {
	n := len(b)
	x := make([]float64, n)

	for i := 0; i < n; i++ {
		sum := b[i]

		for k := 0; k < i; k++ {
			sum -= l[i][k] * x[k]
		}

		x[i] = sum / l[i][i]
	}

	for i := 0; i < n; i++ {
		x[i] = x[i] / d[i][i]
	}

	for i := n - 1; i >= 0; i-- {
		sum := x[i]
		for k := i + 1; k < n; k++ {
			sum -= u1[i][k] * x[k]
		}

		x[i] = sum
	}

	return x
}

func YZ(x [][]float64) ([][]float64, [][]float64) {
	n := len(x)
	y := make([][]float64, n)
	for i := range y {
		y[i] = make([]float64, n)
	}
	z := createMatrix()

	for j := 0; j < n; j++ {
		for i := j; i < n; i++ {
			sum := 0.0
			for k := 0; k < j; k++ {
				sum += y[i][k] * z[k][j]
			}
			y[i][j] = x[i][j] - sum
		}
		for i := 0; i < j; i++ {
			sum := 0.0
			for k := 0; k < i; k++ {
				sum += y[i][k] * z[k][j]
			}
			z[i][j] = (x[i][j] - sum) / y[i][i]
		}
	}

	return y, z
}

func LDU(mu [][]float64) ([][]float64, [][]float64) {
	d := createMatrix()

	for i := 0; i < len(mu); i++ {
		d[i][i] = mu[i][i]
	}

	mu1 := createMatrix()

	for i := 0; i < len(mu); i++ {
		for j := 0; j < len(mu); j++ {
			mu1[i][j] = mu[i][j] / d[i][i]
		}
	}

	return d, mu1
}

func Solve(m [][]float64, ml [][]float64) ([][]float64, [][]float64, [][]float64) {

	mu := copyMatrix(m)

	e := createMatrix()
	einv := createMatrix()

	for j := 0; j < len(mu); j++ {
		for i := j + 1; i < len(mu); i++ {
			if mu[i][j] != 0 {
				factor := mu[i][j] / mu[j][j]

				for k := j; k < len(mu); k++ {
					mu[i][k] = mu[i][k] - factor*mu[j][k]
				}

				e[i][j] = -factor
				einv[i][j] = factor
				ml[i][j] = factor
			}
		}
	}

	return mu, e, einv
}

func main() {
	ma := [][]float64{
		{1, 4, 0},
		{0, 1, 0},
		{0, 3, 1},
	}

	vb := []float64{9, 2, 7}

	ml := createMatrix()

	mu, e, einv := Solve(ma, ml)

	fmt.Println("Пункт 1 \n")

	fmt.Println("матрица A\n=============")
	for str := range ma {
		fmt.Println(ma[str])
	}

	fmt.Println("матрица U\n=============")
	for str := range mu {
		fmt.Println(mu[str])
	}

	fmt.Println("матрица L\n=============")
	for str := range ml {
		fmt.Println(ml[str])
	}

	fmt.Println("матрица E\n=============")
	for str := range e {
		fmt.Println(e[str])
	}

	fmt.Println("матрица E^-1\n=============")
	for str := range einv {
		fmt.Println(einv[str])
	}

	mulLU := mul(ml, mu)
	mulEA := mul(e, ma)

	fmt.Println("матрица L * U = A\n=============")
	for str := range mulLU {
		fmt.Println(mulLU[str])
	}

	fmt.Println("матрица E * A = U \n=============")
	for str := range mulEA {
		fmt.Println(mulEA[str])
	}

	d, mu1 := LDU(mu)

	fmt.Println("матрица D\n=============")
	for str := range d {
		fmt.Println(d[str])
	}

	fmt.Println("матрица U1\n=============")
	for str := range mu1 {
		fmt.Println(mu1[str])
	}

	mulDU1 := mul(d, mu1)
	fmt.Println("матрица D * U1 = U\n=============")
	for str := range mulDU1 {
		fmt.Println(mulDU1[str])
	}

	mulLD := mul(ml, d)
	mulLDU1 := mul(mulLD, mu1)
	fmt.Println("матрица L * D * U1 = A \n=============")
	for str := range mulLDU1 {
		fmt.Println(mulLDU1[str])
	}

	fmt.Println("Пункт 2 \n")

	x := SolveLU(ml, mu, vb)
	fmt.Println("Решение системы уравнений Ax = b\n=============")
	for i := range x {
		fmt.Println(x[i])
	}

	x1 := SolveLDU(ml, d, mu1, vb)
	fmt.Println("Решение системы уравнений Ax = b с использованием метода LDU\n=============")
	for i := range x1 {
		fmt.Println(x1[i])
	}

	fmt.Println("Пункт 3 \n")

	my, mz := YZ(ma)

	fmt.Println("матрица Y\n=============")
	for str := range my {
		fmt.Println(my[str])
	}

	fmt.Println("матрица Z\n=============")
	for str := range mz {
		fmt.Println(mz[str])
	}

	mulYZ := mul(my, mz)
	fmt.Println("матрица Y * Z = X (=A)\n=============")
	for str := range mulYZ {
		fmt.Println(mulYZ[str])
	}

	fmt.Println("проверка: Y == L*D\n=============")
	for str := range mulLD {
		fmt.Println(mulLD[str])
	}

	fmt.Println("проверка: Z == U1\n=============")
	for str := range mu1 {
		fmt.Println(mu1[str])
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

func createMatrix() [][]float64 {
	return [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func copyMatrix(m1 [][]float64) [][]float64 {
	m := make([][]float64, len(m1))

	for i := range m1 {
		m[i] = make([]float64, len(m1[i]))
		copy(m[i], m1[i])
	}

	return m
}
