package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// вывод матрицы в консоль
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

// вычисление матрицы L по исходной матрице А
func calculateL(A *mat.Dense) (L *mat.Dense) {
	n, _ := A.Dims()
	res := mat.NewDense(n, n, nil)
	res.Set(0, 0, math.Sqrt(A.At(0, 0)))
	var el, sum float64
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sum = 0
			if j != 0 {
				for k := 0; k < j; k++ {
					sum += res.At(i, k) * res.At(j, k)
				}
			}
			if i == j {
				el = math.Sqrt(A.At(i, j) - sum)
			} else {
				el = (A.At(i, j) - sum) / res.At(j, j)
			}
			res.Set(i, j, el)
		}
	}
	return res
}

// проверка разложения: L*L^T == A
func check(A, L mat.Dense) (ok bool, res *mat.Dense) {
	n, _ := A.Dims()
	res = mat.NewDense(n, n, nil)
	L1 := L.T()
	res.Mul(&L, L1)
	return mat.EqualApprox(&A, res, 1e-12), res
}

// решение системы методом Холецкого
func sqrtSolve(L, B mat.Dense) (x mat.Dense) {
	y := matSolve(&L, &B)
	x = matSolve(L.T(), &y)
	return x
}

// ____________________________________________________
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

// Транспонирование матрицы в формате Dense
func TDense(A mat.Dense) (res *mat.Dense) {
	n, _ := A.Dims()
	re := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			re.Set(i, j, A.At(j, i))
		}
	}
	return re
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
