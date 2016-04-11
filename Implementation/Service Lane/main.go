package main

import "fmt"

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
	}
}
