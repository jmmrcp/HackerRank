package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		t uint // Numero de casos
		n int  // valor
		i uint // variable para el bucle
	)

	fmt.Scanf("%d", &t) //
	for i = 0; i < t; i++ {
		fmt.Scanf("%d", &n) //
		x := (int(n / 3)) * 3
		for true {
			y := n - x
			if (int(y/5))*5 != y {
				if x == 0 {
					fmt.Println("-1")
					break
				}
				x = (int((x - 1) / 3)) * 3
			} else {
				fmt.Println(strings.Repeat("5", int(x)) + strings.Repeat("3", int(y)))
				break
			}
		}

	}
}
