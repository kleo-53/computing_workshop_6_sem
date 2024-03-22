package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gonum.org/v1/gonum/mat"
)

// вычисление погрешности решений исходной и варьированной матрицы
func printDeviation(x, xv mat.Dense) (dev mat.Dense) {
	dev.Sub(&x, &xv)
	fmt.Println("Погрешность решения:")
	dev = matAbs(dev)
	matPrint(&dev)
	return dev
}

// вывод таблицы на экран
func printTable(data [][]string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Варьирование", "Спектральный", "Объемный", "Угловой"})
	for _, v := range data {
		t.Append(v)
	}
	t.SetRowLine(true)
	t.Render()
}

// вычисление критериев и погрешности решения исходной и варьированной матрицы по левой и правой части уравнения
func calculationsByMatrix(A, varyM, b mat.Dense) {
	a, _ := A.Dims()
	fmt.Println("Исходная матрица размера", a)
	matPrint(&A)
	fmt.Println("")
	fmt.Println("Правая часть уравнения:")
	matPrint(&b)
	fmt.Println("")
	var Av, bv mat.Dense
	Av.Add(&A, &varyM)
	bv.Add(&b, (&varyM).ColView(0))
	fmt.Println("Решение уравнения:")
	X := matSolve(&A, &b)
	matPrint(&X)
	fmt.Println("")
	fmt.Println("Решение варьированного уравнения:")
	Xv := matSolve(&Av, &bv)
	matPrint(&Xv)
	fmt.Println("")
	printDeviation(X, Xv)
	fmt.Println("Критерии: ")
	got := [][]string{
		{fmt.Sprint(0), fmt.Sprint(spectralCriterion(A)), fmt.Sprint(volumeCriterion(A)), fmt.Sprint(angularCriterion(A))},
		{fmt.Sprint(varyM.At(0, 0)), fmt.Sprint(spectralCriterion(Av)), fmt.Sprint(volumeCriterion(Av)), fmt.Sprint(angularCriterion(Av))},
	}
	printTable(got)
}

func main() {
	v := []float64{1, 0.99, 0.99, 0.98}
	A := mat.NewDense(2, 2, v)
	b := mat.NewDense(2, 1, []float64{1.99, 1.97})
	vary := (1e-2 - 1e-10)
	calculationsByMatrix(*A, *mat.NewDense(2, 2, ([]float64{vary, vary, vary, vary})), *b)
}
