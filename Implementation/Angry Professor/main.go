package main

import "fmt"

func main() {
	var (
		t uint // Numero de casos
		k int // Requeridos
		n uint // Numero de alumnos
		a int  // llegadas
		i uint
		j uint
	)

	fmt.Scanf("%d", &t) // 2
	for i = 0; i < t; i++ {
		fmt.Scanf("%d", &n) // 4
		fmt.Scanf("%d", &k) // 3
		for j = 0; j < n; j++ {
			fmt.Scanf("%d", &a) // -1 -3 4 2
			if a <= 0 {
				k--
			}
		}
		if k > 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
