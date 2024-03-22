package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// вывод матрицы на экран
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

// нахождение решения СЛАУ вида Ax=b, возвращает вектор X
func matSolve(A mat.Matrix, b mat.Matrix) (X mat.Dense) {
	err := X.Solve(A, b)
	if err != nil {
		fmt.Printf("Ошибка при решении СЛАУ: %v", err)
		return
	}
	return X
}

// нахождение модуля матрицы
func matAbs(A mat.Dense) (absA mat.Dense) {
	rows, cols := A.Dims()
	absA = *mat.DenseCopyOf(&A)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			el := A.At(i, j)
			if el < 0 {
				absA.Set(i, j, -el)
			}
		}
	}
	return absA
}

// вычисление спектрального критерия по матрице
func spectralCriterion(A mat.Dense) (res float64) {
	var A1 mat.Dense
	err := A1.Inverse(&A)
	if err != nil {
		fmt.Printf("A is not invertible: %v", err)
		return
	}
	res = A.Norm(2) * A1.Norm(2)
	return res
}

// вычисление объемного критерия по матрице
func volumeCriterion(A mat.Dense) (res float64) {
	rows, cols := A.Dims()
	res = 1.0
	cur := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			el := A.At(i, j)
			cur += el * el
		}
		res *= math.Sqrt(cur)
		cur = 0
	}
	res = res / math.Abs(mat.Det(&A))
	return res
}

// вычисление углового критерия по матрице
func angularCriterion(A mat.Dense) (res float64) {
	var C mat.Dense
	err := C.Inverse(&A)
	if err != nil {
		fmt.Printf("A is not invertible: %v", err)
		return
	}
	rows, _ := A.Dims()
	for i := 0; i < rows; i++ {
		res = math.Max(res, math.Abs(mat.Norm(A.RowView(i), 2))*math.Abs(mat.Norm(C.ColView(i), 2)))
	}
	return res
}
