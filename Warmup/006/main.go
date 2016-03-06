package main

import "fmt"

func main() {
	var (
		a uint32 // altura de la escalera > 0
		i uint32
		j uint32
	)

	fmt.Scan(&a)
	for i = 1; i < a+1; i++ {
		for j = 0; j < a; j++ {
			if j < a-i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
