package main

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestIllConditionedMatrix(t *testing.T) {
	cases := []struct {
		in   []float64
		want bool
	}{
		{[]float64{1.0001, 0.9999, 0.9998, 0.9999, 1.0002, 0.9997, 0.9998, 0.9997, 1.0003}, true},
		{[]float64{2.001, 0.999, 0.998, 0.999, 3.002, 0.997, 0.998, 0.997, 4.003}, true},
		{[]float64{10}, true},
		{[]float64{2, 0, 0, 0, 0, 9, 0, 0, 0, 0, 5, 0, 0, 0, 0, 4}, true},
		{[]float64{10, 1, 2, 3, 4, 5, 1, 20, 6, 7, 8, 9, 2, 6, 30, 10, 11, 12, 3, 7, 10, 40, 13, 14, 4, 8, 11, 13, 50, 15, 5, 9, 12, 14, 15, 60}, true},
	}
	for _, c := range cases {
		n := math.Sqrt(float64(len(c.in)))
		A := mat.NewDense(int(n), int(n), c.in)
		fmt.Println("Матрица:")
		matPrint(A)
		fmt.Println("")
		fmt.Printf("Спектральный критерий обусловленности матрицы: %f\n\n", spectralCriterion(*A))
		L := calculateL(A)
		got, LLT := check(*A, *L)
		if got != c.want {
			t.Errorf("get %q, want %q", fmt.Sprint(c.in), fmt.Sprint(c.want))
		} else {
			fmt.Println("L*L^T =")
			matPrint(LLT)
		}
		fmt.Println("_______________________________________________")
	}
}
