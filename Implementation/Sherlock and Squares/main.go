package main

import (
	"fmt"
	"math"
)

func isSquare(x int) bool {
	n := math.Sqrt(float64(x))
	if math.Ceil(n) == math.Floor(n) {
		return true
	}
	return false
}

func answer(x int, y int) {
	var i uint32
	for ; x <= y; x++ {
		if isSquare(x) {
			i++
		}
	}
	fmt.Println(i)
}

func main() {
	var (
		t  uint // Numero de casos
		n1 int  // valor
		n2 int  // valor
		i  uint // variable para el bucle
	)

	fmt.Scanf("%d", &t)
	for i = 0; i < t; i++ {
		fmt.Scanf("%d", &n1)
		fmt.Scanf("%d", &n2)
		answer(n1, n2)
	}
}
