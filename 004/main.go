package main

import "fmt"

func main() {
	var a, i, j uint32
	var b int32

	fmt.Scan(&a)

	arr := make([][]int32, a)
	for i = 0; i < a; i++ {
		arr[i] = make([]int32, a)
	}
	fmt.Printf("%v.\n", arr)
	for i = 0; i < a; i++ {
		for j = 0; j < a; j++ {
			fmt.Scan(&b)
			arr[i][j] = b
		}
	}
	fmt.Printf("%v", arr)
}
