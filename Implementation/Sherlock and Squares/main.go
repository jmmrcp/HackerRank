package main

import (
	"fmt"
	"math"
)

func numCuadrados(x uint32, y uint32) {
	fmt.Println(math.Floor(math.Sqrt(float64(y))) - math.Floor(math.Sqrt(float64(x-1))))
}
func main() {
	var (
		t  uint   // Numero de casos
		n1 uint32 // valor
		n2 uint32 // valor
		i  uint   // variable para el bucle
	)

	fmt.Scanf("%d", &t)
	for i = 0; i < t; i++ {
		fmt.Scanf("%d", &n1)
		fmt.Scanf("%d", &n2)
		numCuadrados(n1, n2)
	}
}
