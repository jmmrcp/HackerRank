package main

import "fmt"

func utree(n int) int {
	tree := 1
	ciclo := 1
	for ciclo <= n {
		if ciclo%2 == 0 {
			tree++
		} else {
			tree *= 2
		}
		ciclo++
	}
	return tree
}

func main() {
	var (
		T int // Numero de casos
		N int // numero de ciclos
		i int
	)
	fmt.Scan(&T)

	for i < T {
		fmt.Scan(&N)
		fmt.Println(utree(N))
		i++
	}
}
