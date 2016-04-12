package main

import (
	"fmt"
	"math"
)

func Multiples(N uint) {
	N = (N - 1)
	n := float64(N)
	var (
		n1   float64
		n2   float64
		n3   float64
		suma float64
	)
	n1 = 3 * (math.Floor(n/3) * (math.Floor(n/3) + 1) / 2)
	n2 = 5 * (math.Floor(n/5) * (math.Floor(n/5) + 1) / 2)
	n3 = 15 * (math.Floor(n/15) * (math.Floor(n/15) + 1) / 2)

	suma = n1 + n2 - n3
	fmt.Println(uint64(suma))
}

func main() {
	var (
		T uint
		N uint
		i uint
	)
	fmt.Scanf("%d", &T)
	for i = 0; i < T; i++ {
		fmt.Scanf("%d", &N)
		Multiples(N)
	}

}
