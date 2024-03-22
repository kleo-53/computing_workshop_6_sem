package main

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestDifferentConditionedMatrix(t *testing.T) {
	cases := []struct {
		inA  []float64
		inB  []float64
		want bool
	}{
		{[]float64{1.0001, 0.9999, 0.9998, 0.9999, 1.0002, 0.9997, 0.9998, 0.9997, 1.0003}, []float64{1, 2, 3}, true},
		{[]float64{2.001, 0.999, 0.998, 0.999, 3.002, 0.997, 0.998, 0.997, 4.003}, []float64{4.1, 0.5, 2.6}, true},
		{[]float64{10}, []float64{2}, true},
		{[]float64{2, 0, 0, 0, 0, 9, 0, 0, 0, 0, 5, 0, 0, 0, 0, 4}, []float64{4, 4.5, 5, 1}, true},
		{[]float64{10, 1, 2, 3, 4, 5, 1, 20, 6, 7, 8, 9, 2, 6, 30, 10, 11, 12, 3, 7, 10, 40, 13, 14, 4, 8, 11, 13, 50, 15, 5, 9, 12, 14, 15, 60}, []float64{9, 1, 3, 8, 12, 6}, true},
	}
	for _, c := range cases {
		n := math.Sqrt(float64(len(c.inA)))
		A := mat.NewDense(int(n), int(n), c.inA)
		B := mat.NewDense(int(n), 1, c.inB)
		fmt.Println("Матрица:")
		matPrint(A)
		fmt.Println("")
		fmt.Printf("Спектральный критерий обусловленности матрицы: %f\n\n", spectralCriterion(*A))
		L := calculateL(A)
		got, LLT := check(*A, *L)
		if got != c.want {
			t.Errorf("Полученные матрицы не равны")
		} else {
			fmt.Println("L*L^T =")
			matPrint(LLT)
		}
		fmt.Println("\nРешение системы с помощью библиотеки:")
		x, xH, eps := matSolve(A, B), sqrtSolve(*L, *B), 1e-10
		matPrint(&x)
		fmt.Println("\nРешение системы c помощью разложения Холецкого")
		matPrint(&xH)
		if mat.EqualApprox(&x, &xH, eps) {
			fmt.Printf("\nОтветы совпадают с точностью %s\n", fmt.Sprint(eps))
		} else {
			fmt.Println("Ответы не совпадают")
			t.Errorf("Ответы не совпадают")
		}
		fmt.Println("_______________________________________________")
	}
}
