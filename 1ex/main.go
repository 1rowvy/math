package main

import (
	"ex1/gauss"
	"ex1/seidel"
	simpleIteration "ex1/simple_iteration"
	"ex1/utils"
	"fmt"
	"math"
)

//A=[[0.6, -0.35, -0.25][0.3, 0.25, 0.43][0.21, -0.45, -0.2]]
//B=[1.83, 0.32, 1.91]

func main() {

	var (
		res       []float64
		iteration int
	)

	matrixA := [][]float64{
		{0.6, -0.35, -0.25},
		{0.3, 0.25, 0.43},
		{0.21, -0.45, -0.2},
	}

	vectorB := []float64{1.83, 0.32, 1.91}

	eps := math.Pow10(-2)

	fmt.Println("Метод простой итерации")
	fmt.Println("=======================")

	newDominantA, newDominantB := func() ([][]float64, []float64) {
		newA := [][]float64{
			{matrixA[0][0] + matrixA[1][0], matrixA[0][1] + matrixA[1][1], matrixA[0][2] + matrixA[1][2]},
			matrixA[2],
			{2*matrixA[1][0] - matrixA[0][0], 2*matrixA[1][1] - matrixA[0][1], 2*matrixA[1][2] - matrixA[0][2]},
		}

		newB := []float64{
			vectorB[0] + vectorB[1],
			vectorB[2],
			2*vectorB[1] - vectorB[0],
		}

		return newA, newB
	}()

	form := utils.ToForm(newDominantA, newDominantB)

	ok, r1, r2, r3, alpha := simpleIteration.CheckConvergence(form)

	fmt.Printf("ρ1 (по строкам)  = %.7f\n", r1)
	fmt.Printf("ρ2 (по столбцам) = %.7f\n", r2)
	fmt.Printf("ρ3 (евклидова)   = %.7f\n", r3)

	if ok {
		fmt.Println("метод сойдётся")
	} else {
		fmt.Println("сходимость не гарантирована")
	}

	target := simpleIteration.CalcTarget(alpha, eps)

	fmt.Println("Решение | Итерация")
	res, iteration = simpleIteration.Solve(form, target)

	fmt.Printf("%f | %d\n", res, iteration)

	fmt.Println("Метод Зейделя")
	fmt.Println("=======================")

	fmt.Println("Решение | Итерация")
	res, iteration = seidel.Solve(form, target)

	fmt.Printf("%f | %d\n", res, iteration)

	fmt.Println("Метод Гаусса")
	fmt.Println("=======================")

	fmt.Println("Решение")
	res = gauss.Solve(matrixA, vectorB)

	fmt.Printf("%f \n", res)
}
