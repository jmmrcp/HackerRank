package main

import "fmt"

func promo(x, y uint) uint {
	if x >= y {
		var a, b, c uint
		a = x / y
		b = x % y
		c = 0
		if a+b >= y {
			c = promo(a+b, y)
		} else if a+b < y {
			return a
		}
		return a + c
	}
	return 0
}

func main() {
	var (
		T uint // Numero de casos
		N uint // Dinero
		C uint // Precio
		M uint // Promo
		i uint
	)

	fmt.Scanf("%d", &T)
	for i = 0; i < T; i++ {
		fmt.Scanf("%d", &N)
		fmt.Scanf("%d", &C)
		fmt.Scanf("%d", &M)

		c := N / C

		fmt.Println(c + promo(c, M))
	}
}
