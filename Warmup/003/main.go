package main

import "fmt"

func main() {
	var a, b, i uint32

	fmt.Scan(&a)

	arr := make([]uint32, a)

	for i = 0; i < a; i++ {
		fmt.Scan(&b)
		arr[i] = b
	}
	fmt.Printf("%v", arr)
}
