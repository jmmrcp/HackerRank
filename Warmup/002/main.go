package main

import "fmt"

func main() {
	var a, b, res, i uint32
	fmt.Scan(&a)

	for i = 0; i < a; i++ {
		fmt.Scan(&b)
		res += b
	}
	fmt.Printf("%v", res)
}
