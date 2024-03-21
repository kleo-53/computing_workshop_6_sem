package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gonum.org/v1/gonum/mat"
)

// подсчет критериев и вывод таблицы в консоль
func printTable(A, L, LT mat.Dense) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Матрица", "Спектральный", "Объемный", "Угловой"})
	data := [][]string{
		{"A", fmt.Sprint(spectralCriterion(A)), fmt.Sprint(volumeCriterion(A)), fmt.Sprint(angularCriterion(A))},
		{"L", fmt.Sprint(spectralCriterion(L)), fmt.Sprint(volumeCriterion(L)), fmt.Sprint(angularCriterion(L))},
		{"L^T", fmt.Sprint(spectralCriterion(LT)), fmt.Sprint(volumeCriterion(LT)), fmt.Sprint(angularCriterion(LT))},
	}
	for _, v := range data {
		t.Append(v)
	}
	t.SetRowLine(true)
	t.Render()
}

func main() {
	A := mat.NewDense(4, 4, []float64{5, 2, 3, 4, 2, 4, 2, 2, 3, 2, 8, 2, 4, 2, 2, 9})
	L := calculateL(A)
	fmt.Println("Матрица L:")
	matPrint(L)
	// matPrint((*TDense(*L)).T().T())
	fmt.Println("")
	fmt.Println("Исходная матрица:")
	matPrint(A)
	fmt.Println("")
	ok, newA := check(*A, *L)
	fmt.Println("Вычисленная L*L^T:")
	matPrint(newA)
	fmt.Println("")
	if ok {
		fmt.Println("A = L*L^T!")
	} else {
		fmt.Println(":(")
	}
	printTable(*A, *L, *TDense(*L))
}
