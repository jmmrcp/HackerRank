package main

import "fmt"

func nundiv(n uint64) uint {
	var (
		s    []uint64
		dig  uint64
		cont uint
		i    int
		num  uint64 = n
	)
	for n > 0 {
		dig = n % 10
		if dig > 0 {
			s = append(s, dig)
		}
		n /= 10
	}
	for i = 0; i < len(s); i++ {
		if num%s[i] == 0 {
			cont++
		}
	}
	return cont
}

func main() {
	var (
		T uint   // Numero de Casos
		N uint64 // Numero
		i uint
	)
	fmt.Scan(&T)
	for i = 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Println(nundiv(N))
	}
}
