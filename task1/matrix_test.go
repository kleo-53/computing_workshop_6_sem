package main

import (
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

// вычисления для рандомной матрицы Гильберта порядка от 2 до 9
func TestHilbMatrix(t *testing.T) {
	a := rand.Intn(8) + 2
	A, varyMatrix, b := mat.NewDense(a, a, nil), mat.NewDense(a, a, nil), mat.NewDense(a, 1, nil)
	vary := (1e-2 - 1e-10)
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			A.Set(i, j, 1.0/(1.0+float64(i)+float64(j)))
			varyMatrix.Set(i, j, vary)
		}
		b.Set(i, 0, rand.Float64()*50)
	}
	calculationsByMatrix(*A, *varyMatrix, *b)
}

// вычисления для трехдиагональной матрицы порядка от 3 до 10
func TestDiagMatrix(t *testing.T) {
	a := rand.Intn(8) + 3
	A, varyMatrix, b := mat.NewDense(a, a, nil), mat.NewDense(a, a, nil), mat.NewDense(a, 1, nil)
	vary := (1e-2 - 1e-10)
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			varyMatrix.Set(i, j, vary)
		}
		if i > 0 {
			A.Set(i, i-1, -1.0)
		}
		if i < a-1 {
			A.Set(i, i+1, -1.0)
		}
		A.Set(i, i, 2.0)
		b.Set(i, 0, rand.Float64()*50)
	}
	calculationsByMatrix(*A, *varyMatrix, *b)
}

// вычисления для случайной матрицы порядка от 2 до 8
func TestRandomMatrix(t *testing.T) {
	a := rand.Intn(7) + 2
	A, varyMatrix, b := mat.NewDense(a, a, nil), mat.NewDense(a, a, nil), mat.NewDense(a, 1, nil)
	vary := (1e-2 - 1e-10)
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			A.Set(i, j, rand.Float64()*10)
			varyMatrix.Set(i, j, vary)
		}
		b.Set(i, 0, rand.Float64()*10)
	}
	calculationsByMatrix(*A, *varyMatrix, *b)
}
