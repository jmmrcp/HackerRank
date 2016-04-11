package main

import "fmt"

func chocolate(N, C, M uint) {
	n := N / C
	fmt.Printf("He can buy %d chocolates with %d and excange the %d wrapper. Total: %d.\n", N/C, N, M, n)
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

		chocolate(N, C, M)
	}
}
